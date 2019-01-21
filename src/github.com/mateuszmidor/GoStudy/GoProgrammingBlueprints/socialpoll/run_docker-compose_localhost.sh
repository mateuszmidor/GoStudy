#!/bin/bash 

# nsqlookup will return HOST_IP address for nsqd.
# This is needed not to return not-reachable from outside docker container ip
export HOST_IP="localhost"
docker-compose up