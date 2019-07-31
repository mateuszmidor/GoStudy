package mykafka

// KafkaAdvertisedListeners holds addresses used in docker-compose_up.yaml for kafka brokers
var KafkaAdvertisedListeners = []string{"172.17.0.1:19092", "172.17.0.1:29092", "172.17.0.1:39092"} // docker static ip
