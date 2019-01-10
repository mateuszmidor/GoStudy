package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"spconfig"

	nsq "github.com/bitly/go-nsq"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Session

func dialdb() error {
	var err error
	config := spconfig.GetConfig()
	log.Println("dialing mongodb: ", config.MongoDbAddress)
	db, err = mgo.Dial(config.MongoDbAddress)
	return err
}

func closedb() {
	db.Close()
	log.Println("MongoDB connection closed")
}

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func publishVotes(votes <-chan string) <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, err := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	exitOnError(err)
	go func() {
		for vote := range votes {
			err := pub.Publish("votes", []byte(vote))
			exitOnError(err)
		}
		log.Println("Producer: stopping...")
		pub.Stop()
		log.Println("Producer: stopped")
		stopchan <- struct{}{}
	}()
	return stopchan
}

func main() {
	var stoplock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		stoplock.Lock()
		stop = true
		stoplock.Unlock()
		log.Println("Stopping...")
		stopChan <- struct{}{}
		closeConn()
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	if err := dialdb(); err != nil {
		log.Fatalln("Couldnt connect to MongoDB: ", err)
	}
	defer closedb()

	votes := make(chan string)
	publisherStoppedChan := publishVotes(votes)
	twitterStoppedChan := startTwitterStream(stopChan, votes)
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			closeConn()
			stoplock.Lock()
			if stop {
				stoplock.Unlock()
				return
			}
			stoplock.Unlock()
		}
	}()
	<-twitterStoppedChan
	close(votes)
	<-publisherStoppedChan
}
