# detect race conditions

## Run with race detector

```sh
go run -race .
```


## Build with race detector

```sh
go build -race .
./race
```


## Test with race detector

```sh
go test -race ./...
```

## Result

```text
==================
WARNING: DATA RACE
Read at 0x0000005c6a40 by goroutine 8:
  main.incrementCounter()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:19 +0x2c

Previous write at 0x0000005c6a40 by goroutine 7:
  main.incrementCounter()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:19 +0x44

Goroutine 8 (running) created at:
  main.main()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:12 +0x33

Goroutine 7 (finished) created at:
  main.main()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:11 +0x27
==================
==================
WARNING: DATA RACE
Write at 0x0000005c6a40 by goroutine 8:
  main.incrementCounter()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:19 +0x44

Previous write at 0x0000005c6a40 by goroutine 7:
  main.incrementCounter()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:19 +0x44

Goroutine 8 (running) created at:
  main.main()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:12 +0x33

Goroutine 7 (finished) created at:
  main.main()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:11 +0x27
==================
==================
WARNING: DATA RACE
Read at 0x0000005c6a40 by main goroutine:
  main.main()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:14 +0x4f

Previous write at 0x0000005c6a40 by goroutine 8:
  main.incrementCounter()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:19 +0x44

Goroutine 8 (finished) created at:
  main.main()
      /home/user/SoftwareDevelopment/GoStudy/race/main.go:12 +0x33
==================
20000
Found 3 data race(s)
exit status 66
```
