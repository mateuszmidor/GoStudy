#!/bin/bash

trap "killall main; killall hw_adapter; killall ui_adapter; killall tuner_adapter; exit 0" SIGINT SIGTERM

cd ui
go run . &
cd ../tuner
go run . &
cd ../hw
go run . &

while true; do sleep 1; done
