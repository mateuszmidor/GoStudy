# gomock example

Doc:  
<https://godoc.org/github.com/golang/mock/gomock>

Tutorial; gomock.Any(), gomock.InOrder() and other useful stuff:  
<https://blog.codecentric.de/en/2017/08/gomock-tutorial/>

## Install

```bash
go get -u github.com/golang/mock/gomock
go get -u github.com/golang/mock/mockgen
```

## Interface of calculator to be mocked (notice the //go:generate)

```go
//go:generate mockgen -source=$GOFILE -destination=$PWD/mocks/${GOFILE} -package=mocks
package calculator

type Calculator interface {
	Add(a, b int) int
	Mul(a, b int) int
}
```

## Generate mocks - reflect mode (needs the //go:generate)

```bash
cd src/
go generate -v ./...
```

## Generate mocks - source mode (doesnt need the //go:generate)

```bash
cd src/
mockgen -source=calculator.go -destination=mocks/calculator.go -package=mocks
```