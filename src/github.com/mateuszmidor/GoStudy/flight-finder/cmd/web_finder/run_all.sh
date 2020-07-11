#!/usr/bin/env bash

trap tearDown SIGINT

function tearDown() {
    pkill -f web_finder
    exit 0
}

go run . &
sleep 1
firefox 'http://localhost:8080/find?from=krk&to=gdn'

while true; do sleep 1; done


