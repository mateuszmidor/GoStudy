package mykafka

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"retry"
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

func ReadMessageWithRetry5(reader *kafka.Reader) (kafka.Message, bool) {
	result := retry.UntilSuccessOr5Failures("reading message", reader.ReadMessage, context.Background())
	if result[1].IsNil() {
		return result[0].Interface().(kafka.Message), true
	}
	return kafka.Message{}, false
}

func DecodeMessageOrLog(r io.Reader, v interface{}) bool {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		fmt.Printf("error while decoding json %s\n", err.Error())
		return false
	}
	return true
}
