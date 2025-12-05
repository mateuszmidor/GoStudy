# modular-monolith

Modules
- sawmill (local module)
- ropeworks (local module)
- sailworks (local module and grpc service)

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

## Run

```sh
    make run

2025/12/05 21:18:25 running sailworks svc at :9000
2025/12/05 21:18:25 NewSawmillLocal client
2025/12/05 21:18:25 NewRopeworksLocal client
2025/12/05 21:18:25 NewSailworksGrpc client: :9000
2025/12/05 21:18:25 received 1 beam
2025/12/05 21:18:25 received 1 beam
2025/12/05 21:18:25 received 1 rope
2025/12/05 21:18:25 received 1 rope
2025/12/05 21:18:25 received 1 rope
2025/12/05 21:18:25 received 1 plank
2025/12/05 21:18:25 received 1 plank
2025/12/05 21:18:25 received 1 plank
2025/12/05 21:18:25 received 1 plank
2025/12/05 21:18:25 received 1 plank
2025/12/05 21:18:25 server received request far new sails: 2
2025/12/05 21:18:25 received 1 sail
2025/12/05 21:18:26 received 1 sail
2025/12/05 21:18:26 received 1 beam
2025/12/05 21:18:26 received 1 rope
2025/12/05 21:18:26 received 1 plank
2025/12/05 21:18:26 received 1 rope
2025/12/05 21:18:26 received 1 rope
2025/12/05 21:18:26 received 1 plank
2025/12/05 21:18:26 received 1 plank
2025/12/05 21:18:26 received 1 plank
2025/12/05 21:18:26 received 1 plank
2025/12/05 21:18:27 received 1 rope
2025/12/05 21:18:27 received 1 plank
2025/12/05 21:18:27 received 1 plank
2025/12/05 21:18:27 received 1 plank
2025/12/05 21:18:27 received 1 plank
2025/12/05 21:18:27 received 1 rope
2025/12/05 21:18:27 received 1 plank
2025/12/05 21:18:27 received 1 rope
2025/12/05 21:18:27 collected 15 planks, 3 beams, 9 ropes, 2 sails
2025/12/05 21:18:27 ship built successfuly
```