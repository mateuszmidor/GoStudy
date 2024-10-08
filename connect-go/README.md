# connect-go example, with go-tags

This demo illustrates the usage of connect-go + gotags in service of gRPC client-server.  
https://connect.build/docs/go/getting-started/

## Install protobuf compiler (protoc) 

1. download the protoc for x64 archive: https://github.com/protocolbufpfers/protobuf/releases 
1. then make it available as terminal command (add it to PATH)
1. test installation:
```bash
    which protoc
    > /home/user/bin/protoc
```

## Install remaining tools

1. make sure `$GOPATH/bin` is on your $PATH
1. run:
```bash
    go get install github.com/bufbuild/buf/cmd/buf
    go install github.com/bufbuild/buf/cmd/buf

    go get google.golang.org/protobuf/cmd/protoc-gen-go
    go install google.golang.org/protobuf/cmd/protoc-gen-go

    go get github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go
    go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go

    go get github.com/srikrsna/protoc-gen-gotag
    go install github.com/srikrsna/protoc-gen-gotag
```

## Initialize or update buf

```bash
buf mod init # only for starting out a new project with "buf" tool
buf mod update # update buf.lock
```

## Run

```bash
make
> server received request: Ping!
> client received response: Pong!
```

## curl it

```bash
curl -H "Content-Type: application/json" -d '{"body": "Ping!"}' localhost:9000/pingpong.PingPong/RpcPing 
> {"body":"Pong!"}
```