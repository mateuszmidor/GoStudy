# go:generate

Example: generate names for enumerators with `stringer` tool.

```bash
#go get -u -a golang.org/x/tools/cmd/stringer 
go install golang.org/x/tools/cmd/stringer 
go generate
go run .
```