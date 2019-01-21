#!/bin/bash

# copy project under $GOPATH/src so it can be "go build" successfully
mkdir -p $GOPATH/src/socialpoll
cd $GOPATH/src/socialpoll
cp -r /home/* .

# install required go packages
echo "go getting envdecode..."
go get github.com/joeshaw/envdecode  
echo "go getting oauth..."
go get github.com/garyburd/go-oauth/oauth  
echo "go getting nsq..."
go get github.com/nsqio/go-nsq
echo "go getting mgo..."
go get gopkg.in/mgo.v2  
echo "getting done."

# build twittervotes, counter, api, web
bash buildall.sh

# setup env variables for aws run
. envconfig/localhost.sh

# run the components
twittervotes/twittervotes &
counter/counter &
cd api
./api &
cd ..
cd web && ./web
