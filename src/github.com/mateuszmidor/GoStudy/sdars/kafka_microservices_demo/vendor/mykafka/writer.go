package mykafka

import (
	"context"
	"encoding/json"
	"io"
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

func Write(writer *kafka.Writer, key string, value []byte) (err error) {
	message := kafka.Message{
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}
	return writer.WriteMessages(context.Background(), message)
}

func EncodeBody(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
