# Go Fuzzing

Test your function against generated input corpus.   
Based on https://go.dev/doc/tutorial/fuzz

## Run

```sh
go test -fuzz=Fuzz -fuzztime 10s # by default, fuzzing goes on until first error is detected
```

Output:

```text
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 34-byte failing input file
fuzz: elapsed: 0s, minimizing
--- FAIL: FuzzReverse (0.10s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:16: Number of runes: orig=1, reversed=2, doubleRev=1
        reverse_test.go:21: Reverse produced invalid UTF-8 string "\x95\xd7"
    
    Failing input written to testdata/fuzz/FuzzReverse/cf43cbb757db554f92c93e6a6321f9853c51dfa19aaf339ba1f07c9bc4dcc163
    To re-run:
    go test -run=FuzzReverse/cf43cbb757db554f92c93e6a6321f9853c51dfa19aaf339ba1f07c9bc4dcc163
FAIL
exit status 1
FAIL    github.com/mateuszmidor/GoStudy/fuzzing 0.108s
```

Input that caused failure is stored in [file](./testdata/fuzz/FuzzReverse/cf43cbb757db554f92c93e6a6321f9853c51dfa19aaf339ba1f07c9bc4dcc163):
```sh
go test fuzz v1
string("ו")
```