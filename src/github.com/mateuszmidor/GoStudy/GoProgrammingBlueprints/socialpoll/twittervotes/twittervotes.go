package main

import (
	"fmt"
	"log"
	"os"

	nsq "github.com/bitly/go-nsq"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("dialing mongodb: localhost")
	db, err = mgo.Dial("localhost")
	return err
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

}
