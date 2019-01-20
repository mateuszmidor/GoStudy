#!/bin/bash

# export Twitter API strings necessary for interacting with Twitter
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "$DIR"/twitter.sh

# export MongoDB & NSQ addresses
export SP_MONGODB_ADDR=ec2-54-93-96-97.eu-central-1.compute.amazonaws.com:27017
export SP_NSQD_ADDR=ec2-54-93-96-97.eu-central-1.compute.amazonaws.com:4150
export SP_NSQLOOKUP_ADDR=ec2-54-93-96-97.eu-central-1.compute.amazonaws.com:4161