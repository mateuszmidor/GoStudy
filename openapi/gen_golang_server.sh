#!/usr/bin/env bash

output="generated_server/"
docker run --rm  -v `pwd`:/local openapitools/openapi-generator-cli generate  -i /local/fridge_api.yaml -g go-server -o /local/$output
sudo chown  -R $USER:$USER  $output # by default the generated code is in user:group root:root