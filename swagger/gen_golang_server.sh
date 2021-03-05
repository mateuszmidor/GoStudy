#!/usr/bin/env bash

output="generated_server"
mkdir -p $output
swagger generate server -t $output -f fridge_api.yaml