#!/bin/bash

# To let Go know the workspace location, do: 
export GOPATH=`pwd` 

# To let shell know go tools, do:
export PATH="`pwd`/bin:$PATH"

#To fix breaking builds eg "go can't find package fmt", do: 
unset GOROOT 
