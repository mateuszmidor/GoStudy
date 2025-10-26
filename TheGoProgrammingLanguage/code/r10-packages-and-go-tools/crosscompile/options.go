// GOARCH=386 go build .  -build 32bit executable. Also: GOARCH=arm, GOARCH=amd64
// GOOS=windows go build . -build windows PE executable. Also: GOOS=darwin, GOOS=linux

package autoinit
 
import "fmt"
 
func init() {
   fmt.Println("Package autoinit initialization")
}

package main // “main” indicates that an executable should be build, not a package

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
