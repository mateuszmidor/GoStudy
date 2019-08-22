package mykafka

import (
	"retry"

	"github.com/segmentio/kafka-go"
)

func newTopic(topic string, numPartitions int, numReplicas int) (string, error) {
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

	return "", lasterror
}

// NewTopicWithRetry5 creates a topic and returns kafka cluster controller address on success
func NewTopicWithRetry5(topic string, numPartitions int, numReplicas int) bool {
	result := retry.UntilSuccessOr5Failures("creating topic", newTopic, topic, numPartitions, numReplicas)
	return result[1].IsNil()
}
