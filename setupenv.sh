#!/bin/bash

# To let Go know the workspace location, do: 
export GOPATH=`pwd` 

#To fix breaking builds eg "go can't find package fmt", do: 
unset GOROOT 
