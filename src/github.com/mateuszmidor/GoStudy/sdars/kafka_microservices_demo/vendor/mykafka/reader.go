package mykafka

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewReader(clientId string, topic string) (w *kafka.Reader) {
	config := kafka.ReaderConfig{
		Brokers:         KafkaAdvertisedListeners,
		GroupID:         clientId,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	return kafka.NewReader(config)
}

func Read(reader *kafka.Reader) (kafka.Message, error) {
	return reader.ReadMessage(context.Background())
}

func DecodeBody(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
