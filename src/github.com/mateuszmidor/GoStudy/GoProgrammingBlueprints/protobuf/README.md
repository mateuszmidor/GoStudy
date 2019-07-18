# protobuf
Remote Procedure Call example using google protobuf

## Install protobuf compiler (protoc)
    https://github.com/protocolbuffers/protobuf/releases (the protoc archive)
    , then make protoc available as terminal command

## Install go-kit, bcrypt and protobuf go components
    go get github.com/go-kit/kit
    go get -u -v golang.org/x/crypto/bcrypt
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    , then make protoc-gen-go from your $GOPATH/bin available as terminal command
    
## Run app:
    ./run_all.sh