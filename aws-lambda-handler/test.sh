#!/usr/bin/env bash

HTTP_ENDPOINT='https://np9ybeaxu9.execute-api.eu-central-1.amazonaws.com/default/hello'
curl -X POST -d '{"name":"Mateusz"}' "$HTTP_ENDPOINT"
echo
