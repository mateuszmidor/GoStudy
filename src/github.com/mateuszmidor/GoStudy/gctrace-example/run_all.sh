#!/usr/bin/env bash

trap tearDown SIGINT

BINARY_NAME="gctrace-example"

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

    echo "OK"
}

function buildExample() {
    stage "Building Go program"

    go build -o $BINARY_NAME .
    echo "OK"
}

function runExample() {
    stage "Running example"

    GOMAXPROCS=3 GODEBUG="gctrace=1" ./$BINARY_NAME
}


function tearDown() {
    pkill -f $BINARY_NAME
    rm $BINARY_NAME
    exit 0
}

checkPrerequsites
buildExample
runExample