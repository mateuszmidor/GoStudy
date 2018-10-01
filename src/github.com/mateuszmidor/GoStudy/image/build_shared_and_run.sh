#!/bin/bash

echo "Building golang libimgutil.so library..."
go build -o libimgutil.so -buildmode=c-shared -ldflags="-s -w"  github.com/mateuszmidor/GoStudy/image/main/

echo "Building c++ test app..."
g++ test.cpp -o test_shared -L . -limgutil

echo "Running test app with ball.jpg param..."
export LD_LIBRARY_PATH=.
./test_shared ball.jpg
