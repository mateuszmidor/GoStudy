# kafka_microservices_demo
Microservices integration over Kafka

Concept:
* HW topic, Tuner topic, UI topic
* MessageID encoded in kafka.Message.Key
* MessageJSON encoded in kafka.Message.Value

## Install golang packages
    go get github.com/segmentio/kafka-go
    go get github.com/segmentio/kafka-go/snappy

## Install docker-compose, pamac(AUR installer), kafkacat(independent kafka CLI consumer)
    sudo pacman -S docker-compose
    sudo pacman -S pamac
    pamac build kafkacat-git

## Test config ready
    ./docker-compos_up.sh       # 3x zookeeper and 3x kafka instances
    ./kafka_make_topic.sh       # create topic "foo" for testing purposes
    ./test_kafka_consumer.sh    # this will read message
    ./test_kafka_producer.sh    # this will send message to consumer

## Run app
    ./run_all.sh    # repeat in case of failure :)

