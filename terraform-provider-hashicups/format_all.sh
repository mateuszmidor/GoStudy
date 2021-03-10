#!/usr/bin/env bash

find . -name '*.go' | xargs gofmt -w