package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"spconfig"

	nsq "github.com/bitly/go-nsq"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

var counts map[string]int
var countsLock sync.Mutex

func doCount(countsLock *sync.Mutex, counts *map[string]int, pollData *mgo.Collection) {
	countsLock.Lock()
	defer countsLock.Unlock()
	if len(*counts) == 0 {
		log.Println("No votes, skipping counting this time")
		return
	}
	log.Println("Updating count database...")
	log.Println(*counts)
	ok := true
	for option, count := range *counts {
		sel := bson.M{"options": bson.M{"$in": []string{option}}}
		up := bson.M{"$inc": bson.M{"results." + option: count}}
		if _, err := pollData.UpdateAll(sel, up); err != nil {
			log.Println("Error updating count database: ", err)
			ok = false
		}
	}
	if ok {
		log.Println("Count database update finish...")
		*counts = nil
	}
}

const updateDuration = 1 * time.Second

func main() {
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	config := spconfig.GetConfig()
	log.Println("Conection to database: ", config.MongoDbAddress)

	db, err := mgo.Dial(config.MongoDbAddress)
	if err != nil {
		fatal(err)
		return
	}
	defer func() {
		log.Println("Disconnecting from database...")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls")
	log.Println("Connection go NSQ messaging system...")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}
	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
		defer countsLock.Unlock()
		if counts == nil {
			counts = make(map[string]int)
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))
	if err := q.ConnectToNSQLookupd("localhost:4161"); err != nil {
		fatal(err)
		return
	}
	ticker := time.NewTicker(updateDuration)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case <-ticker.C:
			doCount(&countsLock, &counts, pollData)
		case <-termChan:
			ticker.Stop()
			q.Stop()
		case <-q.StopChan:
			return
		}
	}
}
