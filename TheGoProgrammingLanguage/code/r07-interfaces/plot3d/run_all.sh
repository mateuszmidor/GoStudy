#!/bin/bash

trap 'exit 0' SIGINT

go run . & 
sleep 0.5s # give the server time to start
firefox localhost:8000/formula

while true; do
sleep 1
done