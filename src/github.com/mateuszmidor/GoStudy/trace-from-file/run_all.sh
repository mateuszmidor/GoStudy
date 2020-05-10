#!/usr/bin/env bash

trap tearDown SIGINT


function stage() {
    GREEN="\e[92m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$GREEN$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"

    command go version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install go compiler to run this example" && exit 1

    echo "OK"
}

function runExample() {
    stage "Running binary and then go tool trace"
    
    go run . > /tmp/gotrace.out
    go tool trace /tmp/gotrace.out

    echo "OK"
}

function keepAlive() {
    stage "CTRL+C to exit"

    while true; do sleep 1; done
}


checkPrerequsites
runExample