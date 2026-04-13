package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lovoo/goka"
	_ "github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/segmentio/kafka-go"
)

var (
	brokers                   = []string{"localhost:9092"}
	inputTopic    goka.Stream = "my-topic-split-in"
	lowerTopic    goka.Stream = "my-topic-split-lower"
	upperTopic    goka.Stream = "my-topic-split-upper"
	group         goka.Group  = "split-processor-group"
	consumerGroup string      = "split-consumer-group"
)

func main() {
	log.SetFlags(log.Ltime)
	createTopics()

	procCtx, procCancel := context.WithCancel(context.Background())

	go runProcessor(procCtx)
	time.Sleep(2 * time.Second)

	producer()
	consumer()

	procCancel()
	time.Sleep(1 * time.Second)
	log.Println("Done")
}

func runProcessor(ctx context.Context) {
	proc := func(ctx goka.Context, msg interface{}) {
		input := msg.(string)
		lower := strings.ToLower(input)
		upper := strings.ToUpper(input)

		ctx.Emit(lowerTopic, ctx.Key(), lower)
		ctx.Emit(upperTopic, ctx.Key(), upper)

		log.Printf("Processor: %q -> (lower: %q, upper: %q)", input, lower, upper)
	}

	groupDef := goka.DefineGroup(group,
		goka.Input(inputTopic, new(codec.String), proc),
		goka.Output(lowerTopic, new(codec.String)),
		goka.Output(upperTopic, new(codec.String)),
	)

	p, err := goka.NewProcessor(brokers, groupDef)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}

	if err := p.Run(ctx); err != nil {
		log.Printf("processor error: %v", err)
	}
}

func createTopics() {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		log.Fatalf("Dial failed: %s\n", err)
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{Topic: string(inputTopic), NumPartitions: 1, ReplicationFactor: 1},
		{Topic: string(lowerTopic), NumPartitions: 1, ReplicationFactor: 1},
		{Topic: string(upperTopic), NumPartitions: 1, ReplicationFactor: 1},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Printf("CreateTopics failed: %s\n", err)
	} else {
		log.Printf("Topics created: %s, %s, %s\n", inputTopic, lowerTopic, upperTopic)
	}
}

func producer() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        string(inputTopic),
		RequiredAcks: -1,
	})
	defer writer.Close()

	for i := range 5 {
		msg := fmt.Sprintf("Hello Kafka %d", i)
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(msg),
		})
		if err != nil {
			log.Printf("WriteMessages failed: %s\n", err)
		} else {
			log.Printf("Produced: %q to %s", msg, inputTopic)
		}
	}

	log.Println("Producer done")
}

func consumer() {
	lowerReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        consumerGroup,
		Topic:          string(lowerTopic),
		StartOffset:    kafka.FirstOffset,
		MinBytes:       1,
		MaxBytes:       1e6,
		MaxWait:        time.Second,
		CommitInterval: time.Second,
	})

	upperReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        consumerGroup,
		Topic:          string(upperTopic),
		StartOffset:    kafka.FirstOffset,
		MinBytes:       1,
		MaxBytes:       1e6,
		MaxWait:        time.Second,
		CommitInterval: time.Second,
	})

	defer lowerReader.Close()
	defer upperReader.Close()

	lowerDone := make(chan struct{})
	upperDone := make(chan struct{})

	go func() {
		for range 5 {
			m, err := lowerReader.FetchMessage(context.Background())
			if err != nil {
				log.Printf("Read from lower failed: %s\n", err)
				continue
			}
			log.Printf("Received lower: %q from %s partition=%d offset=%d\n", string(m.Value), m.Topic, m.Partition, m.Offset)
			lowerReader.CommitMessages(context.Background(), m)
		}
		close(lowerDone)
	}()

	go func() {
		for range 5 {
			m, err := upperReader.FetchMessage(context.Background())
			if err != nil {
				log.Printf("Read from upper failed: %s\n", err)
				continue
			}
			log.Printf("Received upper: %q from %s partition=%d offset=%d\n", string(m.Value), m.Topic, m.Partition, m.Offset)
			upperReader.CommitMessages(context.Background(), m)
		}
		close(upperDone)
	}()

	<-lowerDone
	<-upperDone
	log.Println("Consumer done")
}
