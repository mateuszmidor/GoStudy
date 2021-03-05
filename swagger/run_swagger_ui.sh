#!/usr/bin/env bash

firefox localhost

echo "Please refresh the opened web page"

docker run --rm -p 80:8080 -e SWAGGER_JSON=/local/fridge_api.yaml -v `pwd`:/local swaggerapi/swagger-ui