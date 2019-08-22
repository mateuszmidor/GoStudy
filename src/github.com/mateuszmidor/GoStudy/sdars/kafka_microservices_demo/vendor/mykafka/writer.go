package mykafka

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"retry"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func NewWriter(clientId string, topic string) (w *kafka.Writer) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          KafkaAdvertisedListeners,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	return kafka.NewWriter(config)
}

func WriteMessageWithRetry5(writer *kafka.Writer, key string, value []byte) bool {
	message := kafka.Message{
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	result := retry.UntilSuccessOr5Failures("writing message", writer.WriteMessages, context.Background(), message)
	return result[0].IsNil()
}

func EncodeMessageOrLog(w io.Writer, v interface{}) bool {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Printf("error while encoding json: %s\n", err.Error())
		return false
	}
	return true
}
