# grpc_microservices_demo
Microservices integration over grpc

## Install protobuf compiler (protoc)
    https://github.com/protocolbuffers/protobuf/releases (the protoc for x64 archive)
    , then make protoc available as terminal command (add it to PATH)

## Install protobuf go components
    go get -u google.golang.org/grpc
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    , then make protoc-gen-go from your $GOPATH/bin available as terminal command (add it to PATH)
    

## Run app:
    ./run_all.sh

