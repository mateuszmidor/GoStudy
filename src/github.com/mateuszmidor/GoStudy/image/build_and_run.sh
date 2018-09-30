#!/bin/bash

echo "Building golang libimgutil.so library..."
go build -o libimgutil.so -buildmode=c-shared  github.com/mateuszmidor/GoStudy/image/main/

echo "Building c++ test app..."
g++ test.cpp -o test -L . -limgutil

echo "Running test app with ball.jpg param..."
./test ball.jpg
