#!/bin/bash

function die() {
    printf "Error: $1. Exiting\n"
    exit 1
}

# 1. copy project under $GOPATH/src so it can be "go build" successfully
mkdir -p $GOPATH/src/socialpoll
cd $GOPATH/src/socialpoll
cp -r /home/* . # sdars is mapped as /home (see: docker-compose.yaml)
cd kafka_microservices_demo

# 2. install required go packages
PACKAGES=(github.com/segmentio/kafka-go github.com/segmentio/kafka-go/snappy)
for package in "${PACKAGES[@]}"; do
    echo "go get $package..."
    go get "$package" || die "Failed to go get $package"
    echo "done."
done

# 3. build hw, ui, tuner
COMPONENTS=(hw ui tuner)
for component in "${COMPONENTS[@]}"; do
    echo "Building: $component"
    pushd  "$component" > /dev/null
    go build || die "Failed to build $component"
    popd > /dev/null
    echo "Done."
done

# 4. setup env variables for aws run
#. envconfig.sh || die "Couldnt source envconfig.sh"

# 5. run the components
for component in "${COMPONENTS[@]}"; do
    pushd  "$component" > /dev/null
    ./$component &
    popd > /dev/null
done

# 6. wait then check components running
sleep 3
for component in "${COMPONENTS[@]}"; do
    pushd  "$component" > /dev/null
    pidof "$component" || die "Component $component not running"
    popd > /dev/null
done

# 7. loop until killed
while true; do sleep 1; done