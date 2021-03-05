#!/usr/bin/env bash

output="generated_client"
docker run --rm  -v `pwd`:/local openapitools/openapi-generator-cli generate  -i /local/fridge_api.yaml -g go -o /local/$output
sudo chown  -R $USER:$USER  $output # by default the generated code is in user:group root:root