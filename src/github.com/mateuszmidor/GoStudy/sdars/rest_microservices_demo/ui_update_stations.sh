#!/bin/bash

curl -H 'Content-Type: application/json' -X PUT -d '["jedynka", "dwojka", "trojka"]' localhost:8082/Stations/
