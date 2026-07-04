#!/usr/bin/env bash

trap tearDown SIGINT

IMAGE_NAME="myredis"
REDIS_PORT=6379

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
    
    echo "Done"
}

function runRedis() {
    stage "Running dockerized Redis server"

    docker run \
        -d \
        --rm \
        --name $IMAGE_NAME \
        -p $REDIS_PORT:$REDIS_PORT \
        redis:latest 

    echo "Done"
}

function runExample() {
    stage "Running example"

    go run .

    echo "Done"
}

function keepAlive() {
    stage "CTRL+C to exit"

    while true; do sleep 1; done
}

function tearDown() {
    stage "Stopping..."

    docker stop $IMAGE_NAME

    echo "Done"
    exit 0
}

checkPrerequsites
runRedis
runExample
keepAlive