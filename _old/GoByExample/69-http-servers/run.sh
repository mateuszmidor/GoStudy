#!/usr/bin/env bash

go run . &
sleep 1
curl localhost:8090/hello

pkill -f 69-http-servers