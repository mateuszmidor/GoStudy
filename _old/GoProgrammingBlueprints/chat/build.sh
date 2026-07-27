#!/usr/bin/env bash

# This build.sh is run during AWS deployment, see: " Buildfile"

echo "go getting..."
go get github.com/gorilla/websocket
go get github.com/stretchr/gomniauth
go get github.com/stretchr/testify/mock
go get github.com/clbanning/x2j
go get github.com/ugorji/go/codec
go get gopkg.in/mgo.v2/bson

echo "go building..."
go build -o bin/application
