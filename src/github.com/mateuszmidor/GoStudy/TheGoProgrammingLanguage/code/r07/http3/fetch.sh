#!/bin/bash

telnet <<< $'open 127.0.0.1 8000\nGET /list HTTP/1.1\nHost: 127.0.0.1:8000\n\n' 


#<< 'EOF'
#'GET /list HTTP/1.1' 
#'Host: 127.0.0.1:8000'
#\n
#EOF
