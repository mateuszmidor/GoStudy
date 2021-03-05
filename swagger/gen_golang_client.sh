#!/usr/bin/env bash

output="generated_client"
mkdir -p $output
swagger generate client -t $output -f fridge_api.yaml