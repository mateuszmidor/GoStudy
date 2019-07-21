# kafka_microservices_demo
Microservices integration over Kafka

## Install golang packages
    go get github.com/segmentio/kafka-go
    go get github.com/segmentio/kafka-go/snappy

## Install docker-compose, pamac(AUR installer), kafkacat(independent kafka CLI consumer)
    sudo pacman -S docker-compose
    sudo pacman -S pamac
    pamac build kafkacat-git

## Test config ready
    ./docker-compos_up.sh       # 3x zookeeper and 3x kafka instances
    ./kafka_make_topic.sh       # create topic "foo" for testing purposesf
    ./test_kafka_consumer.sh    # this will read message
    ./test_kafka_producer.sh    # this will send message to consumer

## Run app
    

