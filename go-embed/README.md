# go-embed

This example demonstrates the usage of "embed" Go package together with `//go:embed` directive.  
They together pupulate the variable defined below the directive with contents of indicated file, during application build.

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed welcome.txt
var welcomeMessage string // this variable gets populated with contents of 'welcome.txt' when building the app

func main() {
	fmt.Println(welcomeMessage)
}
```