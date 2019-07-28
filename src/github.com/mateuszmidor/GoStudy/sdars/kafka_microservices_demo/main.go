package main

// BASED ON: https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26
import (
	"fmt"
	"log"
	"time"
	"utils"

	"mykafka"
)

const topic = "IntraProcessTopic"

func readLoop() {
	reader := mykafka.NewReader("reader", topic)
	defer reader.Close()

	for {
		m, err := mykafka.Read(reader)
		if err != nil {
			log.Fatalf("error while receiving message: %s", err.Error())
			continue
		}
		value := m.Value

		if err != nil {
			log.Fatalf("error while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
	}
}

func writeLoop() {
	writer := mykafka.NewWriter("writer", topic)
	defer writer.Close()

	for {
		current := time.Now()
		err := mykafka.Write(writer, nil, []byte("Kafka demo "+current.Format("15:04:05")))
		if err != nil {
			log.Fatal("Kafka write error", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func createTopic(topic string) {
	controlleraddr, lasterror := mykafka.NewTopic(topic, 6, 2)
	if lasterror == nil {
		log.Printf("Created topic '%s' using cluster controller: %s\n", topic, controlleraddr)
		return // success
	}

	log.Fatalf("CreateTopics error %v\n", lasterror)
}

func deleteTopic(topic string) {
}

func main() {
	createTopic(topic)

	go readLoop()
	go writeLoop()

	utils.NewShutdownCondition().Wait()
	deleteTopic(topic)
}
