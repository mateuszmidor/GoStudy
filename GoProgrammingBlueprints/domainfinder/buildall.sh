#!/bin/bash

# Build all the tool binaries
DOMAIN_FINDER='.'
TOOLS=(synonyms sprinkle coolify domanify available "$DOMAIN_FINDER")
for tool in "${TOOLS[@]}"; do
    pushd  "$tool" > /dev/null
    echo "Building: $tool"
    go build
    echo "Done."
    popd > /dev/null
done
