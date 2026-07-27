# Watch for file changes and rerun the app

This is good for testing local servers.

```bash
# install 'reflex' tool, it gets built in $GOPATH/bin"
go get github.com/cespare/reflex

# run 'reflex': watch .go files for changes, and if change is seen - execute 'go run .'
$GOPATH/bin/reflex -r '\.go$' go run .
```
