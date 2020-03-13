# protobuf

Microservices integration using alternatively REST or gRPC through go-kit framework

## Install protobuf compiler (protoc)

    https://github.com/protocolbuffers/protobuf/releases (the protoc archive)
    , then make protoc available as terminal command

## Install go-kit, bcrypt and protobuf go components

    go get github.com/go-kit/kit
    go get golang.org/x/time/rate
    go get -u -v golang.org/x/crypto/bcrypt
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    , then make protoc-gen-go from your $GOPATH/bin available as terminal command

## Generate .go from .proto

    cd <folder with .proto files>
    protoc *.proto --go_out=plugins=grpc:.

## Run app

    ./run_all.sh
