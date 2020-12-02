# //go:nosplit demo

Based on <https://dave.cheney.net/2018/01/08/gos-hidden-pragmas>  
This demo shows that the red zone can be exhausted and stack overflow happen when using //go:nosplit  
Note: this works in go 1.11, later not, eg.

```bash
docker run -it -v `pwd`:/go golang:1.11 go run src/main.go

main.C: nosplit stack overflow
        744     assumed on entry to main.A (nosplit)
        480     after main.A (nosplit) uses 264
        472     on entry to main.B (nosplit)
        208     after main.B (nosplit) uses 264
        200     on entry to main.C (nosplit)
        -64     after main.C (nosplit) uses 264
```
