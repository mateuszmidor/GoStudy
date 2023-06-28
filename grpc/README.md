# grpc

This demo illustrates the usage of gRPC as ping-pong server

## Install protobuf compiler (protoc)

1. download the protoc for x64 archive: https://github.com/protocolbuffers/protobuf/releases 
1. then make it available as terminal command (add it to PATH)
1. test installation:
```bash
    which protoc
    > /home/user/bin/protoc
```

## Install Go plugin

1. make sure `$GOPATH/bin` is on your $PATH
1. run:
```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Run

```bash
make
> server received request: Ping!
> client received response: Pong!
```