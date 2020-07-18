#!/usr/bin/env bash

trap tearDown SIGINT

function tearDown() {
    pkill -f finder_web
    exit 0
}

go run . &
sleep 1
firefox 'http://localhost:9090/'

while true; do sleep 1; done