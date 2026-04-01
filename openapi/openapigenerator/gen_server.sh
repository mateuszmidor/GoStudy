#!/bin/bash
set -e

docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  -i /local/fridge_api.yaml \
  -g go-server \
  -o /local/generated_server \
  --additional-properties=packageName=server,useGinFramework=false,moduleName=github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_server

echo "Server generated in generated_server/"
