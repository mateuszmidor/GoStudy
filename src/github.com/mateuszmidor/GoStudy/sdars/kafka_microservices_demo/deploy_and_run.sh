#!/bin/bash

function die() {
    printf "Error: $1. Exiting\n"
    exit 1
}

# 1. copy project under $GOPATH/src so it can be "go build" successfully
mkdir -p $GOPATH/src/sdars
cd $GOPATH/src/sdars
cp -r /home/* . # sdars is mapped as /home (see: docker-compose.yaml)
cd kafka_microservices_demo

# 2. install required go packages
PACKAGES=(github.com/segmentio/kafka-go github.com/segmentio/kafka-go/snappy)
for package in "${PACKAGES[@]}"; do
    echo "go get $package..."
    go get "$package" || die "Failed to go get $package"
    echo "done."
done

# 3. build component provided as run param (either hw/tuner/ui)
COMPONENT="$1"
echo "Building: $COMPONENT"
cd "$COMPONENT"
go build || die "Failed to build $COMPONENT"
echo "Done."

# 4. run the component
./$COMPONENT
