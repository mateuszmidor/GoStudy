#!/bin/bash

trap 'exit 0' SIGINT

go run . & 
sleep 0.5s # give the server time to start
telnet localhost 8000
