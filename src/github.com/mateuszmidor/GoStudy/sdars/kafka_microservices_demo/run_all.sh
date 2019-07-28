#!/bin/bash

trap "killall main; killall hw_adapter; killall ui_adapter; killall tuner_adapter; exit 1" SIGINT SIGTERM

./docker-compose_up.sh 

exit 0

sleep 10

cd ui
go run . &
cd ../tuner
go run . &
cd ../hw
go run .