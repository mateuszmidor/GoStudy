package mykafka

import (
	"context"
	"encoding/json"
	"fmt"
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

func ReadMessageOrLog(reader *kafka.Reader) (kafka.Message, error) {
	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		fmt.Printf("error while receiving message: %s\n", err.Error())
	}
	return msg, err
}

func DecodeMessageOrLog(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		fmt.Printf("error while decoding json %s\n", err.Error())
	}
	return err
}
