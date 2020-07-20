#!/usr/bin/env bash

go build . # -gcflags=-B .
./finder_cli < test_cases
rm finder_cli