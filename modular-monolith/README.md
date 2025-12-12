# modular-monolith

This project models a shipyard as a modular monolith:
- the monolith is made of departments (implemented as modules)
- every module is implemented as a private(internal) package, but exposes a public API
  - internal guarantees that other modules don't reference the internals but only use the public API
- this imposes initial effort of having private domain objects being translated to public API objects
  -  how to avoid it?

## Modules

- sawmill (extracted as a separate grpc service)
- ropeworks (local module)
- mastworks (local module)
- sailworks (local module)

![Architecture Diagram](./docs/C4_Component.png)


## Prerequisites

GRPC is used for communication with sawmill module, thus GRPC tooling for Go must be installed.

### Install protobuf compiler (protoc)

1. download the protoc for x64 archive: https://github.com/protocolbufpfers/protobuf/releases
1. then make it available as terminal command (add it to PATH)
1. test installation:
```bash
    which protoc
    > /home/user/bin/protoc
```

### Install remaining tools

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
```
```log
  2025/12/12 16:15:26 NewSawmillGrpc client: :9001
  2025/12/12 16:15:26 NewRopeworksLocal client
  2025/12/12 16:15:26 NewMastworksLocal client
  2025/12/12 16:15:26 NewSailworksLocal client
  2025/12/12 16:15:26 Sailworks produced 1 sail
  2025/12/12 16:15:26 Sailworks produced 1 sail
  2025/12/12 16:15:26 Ropeworks produced 1 rope
  2025/12/12 16:15:26 Ropeworks produced 1 rope
  2025/12/12 16:15:26 Ropeworks produced 1 rope
  2025/12/12 16:15:26 SawmillService received GetBeams request: 3
  2025/12/12 16:15:26 Sawmill produced 1 beam
  2025/12/12 16:15:26 Sawmill produced 1 beam
  2025/12/12 16:15:26 Sawmill produced 1 beam
  2025/12/12 16:15:26 Mastworks received 3 beams for making a mast
  2025/12/12 16:15:26 Mastworks produced 1 mast
  2025/12/12 16:15:26 SawmillService received GetBeams request: 3
  2025/12/12 16:15:26 Sawmill produced 1 beam
  2025/12/12 16:15:26 Sawmill produced 1 beam
  2025/12/12 16:15:26 Sawmill produced 1 beam
  2025/12/12 16:15:26 Mastworks received 3 beams for making a mast
  2025/12/12 16:15:27 Sailworks produced 1 sail
  2025/12/12 16:15:27 Sailworks produced 1 sail
  2025/12/12 16:15:27 Ropeworks produced 1 rope
  2025/12/12 16:15:27 Ropeworks produced 1 rope
  2025/12/12 16:15:27 Ropeworks produced 1 rope
  2025/12/12 16:15:27 Mastworks produced 1 mast
  2025/12/12 16:15:28 Ropeworks produced 1 rope
  2025/12/12 16:15:28 Ropeworks produced 1 rope
  2025/12/12 16:15:28 Ropeworks produced 1 rope
  2025/12/12 16:15:28 collected 9 ropes, 2 masts, 4 sails
  2025/12/12 16:15:28 ### ship built successfuly ###
```