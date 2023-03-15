# GO 1.18 workspaces

Based on https://go.dev/doc/tutorial/workspaces  
Replaces external dependency `golang.org/x/example` with it's local, modified copy under ./example using [go.work](./go.work) file.  
This replacement is in force across all modules in workspace - you don't need to add `replace` directive in every module's `go.mod`.  
Seems this is the intended profit - to work with multiple modules easily.  
This feature can be disabled with env variable `GOWORK=off`

## Run it

Thank to workspaces, you can `go run` from the workspace, not necessarily from the module:

```sh
go run example.com/hello
```