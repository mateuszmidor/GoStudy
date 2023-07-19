# errgroup

This example demonstrates how to synchronize goroutines and propagate errors from goroutines using package` golang.org/x/sync/errgroup`.
* https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/
* https://www.youtube.com/watch?v=KGOgEr7tFKE&list=PL7yAAGMOat_F7bOImcjx4ZnCtfyNEqzCy&index=10

## Run

```bash
go run .
```

```
2023/07/19 22:56:57 100/10 = 10
2023/07/19 22:56:57 100/8 = 12.5
2023/07/19 22:56:57 100/9 = 11.111111
2023/07/19 22:56:57 100/6 = 16.666666
2023/07/19 22:56:57 100/5 = 20
2023/07/19 22:56:57 100/7 = 14.285714
2023/07/19 22:56:57 100/3 = 33.333332
2023/07/19 22:56:57 100/2 = 50
2023/07/19 22:56:57 100/1 = 100
2023/07/19 22:56:57 100/4 = 25
2023/07/19 22:56:57 error: division by zero
```