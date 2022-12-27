#!/usr/bin/env bash

trap tearDown SIGINT

function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"

    command docker --version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install docker to run this example" && exit 1

    echo "Done"
}

function runPrometheus() {
    stage "Running Prometheus"

    docker run -d --rm --name=my_prometheus --net=host -v `pwd`:/home  prom/prometheus --config.file="/home/config.yml"

    echo "Find your metric under 'wave' name!!!"
    sleep 3
    firefox localhost:9090/targets # Prometheus dashboard
    
    echo "Done"
}

function runMetricProvider() {
    stage "Running 'wave' metric(a sinus function) provider at port 8080"

    go run .

    echo "Done"
}


function tearDown() {
    stage "Tear down"

    docker stop my_prometheus

    echo "Done"
    exit 0
}

function keepAlive() {
    # keep alive to intercept CTLR+C and run tearDown and kill prometheus
    stage "CTRL+C to exit"

    while true; do sleep 1; done
}

checkPrerequsites
runPrometheus
runMetricProvider
keepAlive
