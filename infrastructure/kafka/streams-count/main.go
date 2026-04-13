package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lovoo/goka"
	_ "github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/segmentio/kafka-go"
)

var (
	brokers                   = []string{"localhost:9092"}
	inputTopic    goka.Stream = "my-topic-count-in"
	outputTopic   goka.Stream = "my-topic-count-out"
	group         goka.Group  = "count-processor-group"
	consumerGroup string      = "count-consumer-group"
)

func main() {
	log.SetFlags(log.Ltime)
	createTopics()

	procCtx, procCancel := context.WithCancel(context.Background())

	// run stream processor first as it only receives input events from the most recent one
	// so if producer was run first the processor would miss the events
	go runProcessor(procCtx)
	time.Sleep(3 * time.Second) // let  the processor initialize

	producer()
	consumer()

	procCancel()
	time.Sleep(12 * time.Second)
	log.Println("Done")
}

// this stream processor is stateful - keeps total num of characters in /tmp/goka database
func runProcessor(ctx context.Context) {
	// Configure replication factor to 1 because the cluster only has 1 broker (default is 2)
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	proc := func(ctx goka.Context, msg interface{}) {
		input := msg.(string)
		var totalCharCount int64
		if val := ctx.Value(); val != nil {
			totalCharCount = val.(int64)
		}
		totalCharCount += int64(len(input))
		ctx.SetValue(totalCharCount)
		output := fmt.Sprintf("%s (%d total chars received)", input, totalCharCount)
		ctx.Emit(outputTopic, ctx.Key(), output)
		log.Printf("Processor: %q -> %q", input, output)
	}

	groupDef := goka.DefineGroup(group,
		goka.Input(inputTopic, new(codec.String), proc),
		goka.Output(outputTopic, new(codec.String)),
		goka.Persist(new(codec.Int64)), // this makes the processor stateful
	)

	p, err := goka.NewProcessor(brokers, groupDef,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
	)
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
		{Topic: string(outputTopic), NumPartitions: 1, ReplicationFactor: 1},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Printf("CreateTopics failed: %s\n", err)
	} else {
		log.Printf("Topics created: %s, %s\n", inputTopic, outputTopic)
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
			Key:   []byte("mykey"),
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
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           brokers,
		GroupID:           consumerGroup,
		StartOffset:       kafka.FirstOffset,
		HeartbeatInterval: time.Second * 3,
		SessionTimeout:    time.Second * 9,
		Topic:             string(outputTopic),
	})
	defer r.Close()

	for range 5 {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			log.Printf("Read failed: %s\n", err)
			continue
		}
		log.Printf("Received: %q from %s partition=%d offset=%d\n", string(m.Value), m.Topic, m.Partition, m.Offset)
		if err := r.CommitMessages(context.Background(), m); err != nil {
			log.Printf("Commit failed: %s\n", err)
		}
	}
	log.Println("Consumer done")
}
