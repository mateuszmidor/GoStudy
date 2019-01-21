#!/bin/bash

COMPONENTS=(twittervotes counter api web)
for component in "${COMPONENTS[@]}"; do
    pushd  "$component" > /dev/null
    echo "Building: $component"
    go build
    echo "Done."
    popd > /dev/null
done