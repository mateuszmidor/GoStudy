# dependency injection example

Simple injector, matches exact type and "inject: yes" tag

## Example

```go
package main

import (
	"fmt"
	"reflect"
)

type person struct {
	FirstName string `inject:"yes"`
	LastName  string `inject:"no"`
	Age       uint8  `inject:"yes"`
	email     string
}

func main() {
	injector := NewInjector()
	injector.Set(reflect.TypeOf(string("")), "Jessica")
	injector.Set(reflect.TypeOf(uint8(0)), uint8(22))

	var p person
	injector.Inject(&p)
	fmt.Printf("%+v", p)
}
```
Output:  
```{FirstName:Jessica LastName: Age:22 email:}```
