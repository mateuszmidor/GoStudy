#!/bin/bash
set -e

docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  -i /local/fridge_api.yaml \
  -g go \
  -o /local/generated_client \
  --additional-properties=packageName=client,moduleName=github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_client

echo "Client generated in generated_client/"
