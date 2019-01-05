#!/bin/bash

TOOLS=(synonyms sprinkle coolify domanify available)
for tool in "${TOOLS[@]}"; do
    pushd "$tool"
    go build
    popd
done
go build 