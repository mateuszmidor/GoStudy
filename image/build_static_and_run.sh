#!/bin/bash

echo "Building golang libimgutil.a library..."
go build -o libimgutil.a -buildmode=c-archive -ldflags="-s -w" github.com/mateuszmidor/GoStudy/image/

echo "Building c++ test app..."
g++ test.cpp -o test_static -static -L . -limgutil -pthread

echo "Running test app with ball.jpg param..."
./test_static ball.jpg
