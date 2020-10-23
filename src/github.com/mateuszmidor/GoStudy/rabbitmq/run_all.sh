#!/usr/bin/env bash

trap tearDown SIGINT

IMAGE_NAME="rabbitmq-server"
SERVICE_PORT=5672
WEB_CONSOLE_PORT=15672
WEB_CONSOLE_URL="localhost:$WEB_CONSOLE_PORT"

function stage() {
    BLUE="\e[36m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BLUE$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"

    command go version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install go compiler to run this example" && exit 1

    command docker version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install docker to run this example" && exit 1

    echo "OK"
}

function buildExample() {
    stage "Building producer and consumer"

    go build -o producer/producer producer/main.go
    go build -o consumer/consumer consumer/main.go
    echo "OK"
}

function runRabbitMQ() {
    stage "Running dockerized RabbitMQ server"

    docker run -d --rm --hostname my-rabbit --name $IMAGE_NAME -p $SERVICE_PORT:$SERVICE_PORT -p $WEB_CONSOLE_PORT:$WEB_CONSOLE_PORT rabbitmq:3-management

    # wait server is up
    waitRabbitReady
    echo "OK"
}

function runProducer() {
    stage "Running producer"

    "./producer/producer"

    echo "OK"
}

function runConsumer() {
    stage "Running consumer"

    "./consumer/consumer"

    echo "OK"
}

function runManagementConsole() {
    stage "Opening management console"

    echo "Login/Pass: guest/guest"
    sleep 5
    firefox $WEB_CONSOLE_URL
}

function waitRabbitReady {
    
    while true; do curl --max-time 1 $WEB_CONSOLE_URL > /dev/null 2>&1; [[ $? == 0 ]] && break; sleep 1; done
}

function keepAlive() {
    stage "CTRL+C to exit"

    while true; do sleep 1; done
}

function tearDown() {
    stage "Stopping..."

    docker stop $IMAGE_NAME
    exit 0
}

checkPrerequsites
buildExample
runRabbitMQ
runProducer
runConsumer
runManagementConsole
keepAlive