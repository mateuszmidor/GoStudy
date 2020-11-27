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

    echo "Done"
}

function pushd () {
    command pushd "$@" > /dev/null # silent pushd
}

function popd () {
    command popd "$@" > /dev/null # silent popd
}

function generateCertificate() {
    stage "Generating TLS certificate"

    if [[ ! -e ../cert/minica/localhost ]]; then
        pushd ../cert/minica 
        go run . --domains localhost
        popd
    fi

    echo "Done"
}

function runServer() {
    stage "Running tls server"

    pushd src/server
    go run . &
    popd

    sleep 3 # give server time to startup

    echo "Done"
}

function runClient() {
    stage "Running tls client"

    pushd src/client
    go run .
    popd

    echo "Done"
}

function keepAlive() {
    stage "CTRL+C to exit"

    while true; do sleep 1; done
}

function tearDown() {
    stage "Stopping..."


    echo "Done"
    exit 0
}

checkPrerequsites
generateCertificate
runServer
runClient
keepAlive