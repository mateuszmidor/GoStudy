#!/bin/bash

# make it USA NYC time
TZ=US/Eastern go run . & 
sleep 1
telnet 127.0.0.1 8000
