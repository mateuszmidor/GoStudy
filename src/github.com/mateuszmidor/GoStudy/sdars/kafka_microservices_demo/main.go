package main

// BASED ON: https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26
import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var writer *kafka.Writer

func Configure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}

func Push(parent context.Context, key, value []byte) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return writer.WriteMessages(parent, message)
}

const myID string = "kafkaGoProducerTest"
const topic = "foo" // IMPORTANT: topic was preconfigured with 2 replications and 4 partitions

var (
	// kafka
	kafkaBrokerUrl     string = "localhost:19092,localhost:29092,localhost:39092"
	kafkaVerbose       bool   = true
	kafkaTopic         string = topic
	kafkaConsumerGroup string = "kafka-consumer-group"
	kafkaClientId      string = "kafkaGoClientId"
)

func readLoop() {
	// flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers in comma separated value")
	// flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	// flag.StringVar(&kafkaTopic, "kafka-topic", "foo", "Kafka topic. Only one topic per worker.")
	// flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	// flag.StringVar(&kafkaClientId, "kafka-client-id", "my-client-id", "Kafka client id")

	// flag.Parse()

	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	brokers := strings.Split(kafkaBrokerUrl, ",")

	// make a new reader that consumes from topic-A
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         kafkaClientId,
		Topic:           kafkaTopic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error while receiving message: %s", err.Error())
			continue
		}

		value := m.Value
		// if m.CompressionCodec == snappy.NewCompressionCodec() {
		// 	_, err = snappy.NewCompressionCodec().Decode(value, m.Value)
		// }

		if err != nil {
			log.Fatalf("error while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
	}
}

func main() {

	go readLoop()

	Configure([]string{"localhost:19092", "localhost:29092", "localhost:39092"}, myID, topic)

	for {

		current := time.Now()
		err := Push(context.Background(), nil, []byte("Kafka demo "+current.Format("15:04:05")))
		if err != nil {
			log.Fatal("Kafka write error", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
