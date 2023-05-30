# fanout/loadbalancer

Inspired by "Concurrency Patterns In Go" <https://youtu.be/YEKjSzIwAdA?t=891>

```go
func fanOut(in []int, out1 chan<- int, out2 chan<- int) {
    for i := range in {
        select {
        case out1 <- i:
        case out2 <- i:
        }
    }
}
```
