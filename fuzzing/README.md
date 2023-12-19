# Go Fuzzing

Test your function against generated input corpus.
Based on https://go.dev/doc/tutorial/fuzz

## Run

```sh
go test -fuzz=Fuzz reverse_test.go -fuzztime 10s # by default, fuzzing goes on until first error is detected
```

Output:

```text
=== RUN   FuzzReverse
fuzz: elapsed: 0s, gathering baseline coverage: 0/13 completed
fuzz: elapsed: 0s, gathering baseline coverage: 13/13 completed, now fuzzing with 16 workers
fuzz: elapsed: 0s, execs: 386 (5902/sec), new interesting: 0 (total: 13)
--- FAIL: FuzzReverse (0.07s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:17: Number of runes: orig=1, reversed=2, doubleRev=1
        reverse_test.go:22: Reverse produced invalid UTF-8 string "\xb0\xc3"

    Failing input written to testdata/fuzz/FuzzReverse/ddf3b8ab3d7347d1
    To re-run:
    go test -run=FuzzReverse/ddf3b8ab3d7347d1
=== NAME
FAIL
exit status 1
FAIL    command-line-arguments  0.075s
```

Input that caused failure is stored under [file](./testdata/fuzz/FuzzReverse/ddf3b8ab3d7347d1):

```
go test fuzz v1
string("ð")
```
, meaning that input "ð" caused the test to fail.