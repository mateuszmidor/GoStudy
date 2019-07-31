package mykafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

// NewTopic creates a topic and returns kafka cluster controller address on success
func NewTopic(topic string, numPartitions int, numReplicas int) (string, error) {
	brokers := KafkaAdvertisedListeners

	var lasterror error

	// find current controler among the brokers
	for _, possibleControllerAddr := range brokers {
		conn, err := kafka.Dial("tcp", possibleControllerAddr)
		if err != nil {
			lasterror = err
			continue
		}

		// fmt.Printf("NewTopic: dialed to %s\n", possibleControllerAddr)
		t := kafka.TopicConfig{
			Topic:              topic,
			NumPartitions:      numPartitions,
			ReplicationFactor:  numReplicas,
			ReplicaAssignments: nil,
			ConfigEntries:      nil,
		}

		lasterror = conn.CreateTopics(t)
		if lasterror == nil {
			// fmt.Printf("NewTopic: created with %s\n", possibleControllerAddr)
			return possibleControllerAddr, nil
		}
	}

	fmt.Printf("NewTopic: %s\n", lasterror.Error())
	return "", lasterror
}
