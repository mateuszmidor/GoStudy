#!/bin/bash 

# this will run docker-compose with input recipe: docker-compose.yaml
export MY_IP=172.17.0.1 # this IP is docker static ip, invisible from outside
docker-compose up --abort-on-container-exit --force-recreate
