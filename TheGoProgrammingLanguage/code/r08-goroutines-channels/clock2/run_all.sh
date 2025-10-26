#!/bin/bash

trap 'exit 0' SIGINT

# make it USA NYC time
TZ=US/Eastern go run . & 
sleep 0.5s
telnet 127.0.0.1 8000
