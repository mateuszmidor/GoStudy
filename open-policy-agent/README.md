# open-policy-agent

This example demonstrates usage of OPA in Go.  
It checks if people are of allowed age for drinking (age >= 18).  

## Policy

```rego
package age_authorization

import future.keywords

default can_drink := false # can_drink will be available as policy evaluation output

can_drink if {
	input.age >= 18
}
```

https://github.com/open-policy-agent/opa

## Run

```bash
go run .
```

```
2023/07/04 22:39:13 [Kira,22   ]: map[can_drink:true]
2023/07/04 22:39:13 [Crystal,17]: map[can_drink:false]
```