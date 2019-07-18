#!/bin/bash

trap "killall main; killall hw_adapter; killall ui_adapter; killall tuner_adapter" SIGINT SIGTERM

# generate go files from proto files
command protoc > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "Generating .go files from .proto..."
    pushd vendor/rpc > /dev/null
    ./generate_proto.sh
    echo "Generating done."
    popd > /dev/null
else
    echo "protoc command not available. Using pregenerated .pb.go files"
fi

# run all services
cd ui
go run . &
cd ../tuner
go run . &
cd ../hw
go run .