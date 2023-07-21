#!/usr/bin/env bash

trap tearDown SIGINT


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

    command docker-compose version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install docker-compose to run this example" && exit 1

    echo "Done"
}

function runCockroachCluster() {
    stage "Running Cockroach cluster"

    docker-compose up &
    sleep 5 # let the instances run; should be done some smarter way
    docker exec node_1 ./cockroach init --insecure
    sleep 5 # let the instances initialize

    echo "Done"
}

function runExample() {
    stage "Running example"

    cd rawsql/
    go run .
    firefox http://localhost:8080

    echo "Done"
}

function keepAlive() {
    stage "CTRL+C to exit"

    while true; do sleep 1; done
}

function tearDown() {
    stage "Stopping..."

    docker-compose down

    echo "Done"
    exit 0
}

checkPrerequsites
runCockroachCluster
runExample
keepAlive