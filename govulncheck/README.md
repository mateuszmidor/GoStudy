# govulncheck

* Check for vulnerabilities in Go modules and packages
* https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck
* https://www.youtube.com/watch?v=BOQfO60gWGM&list=PL7yAAGMOat_GCd12Lrv_evJ3Zhv1dl8B-&index=38

## Install

```sh
go install golang.org/x/vuln/cmd/govulncheck@latest
```
## Run against source code

```sh
govulncheck ./...
```

```
Scanning your code and 49 packages across 1 dependent module for known vulnerabilities...

Vulnerability #1: GO-2023-1840
    Unsafe behavior in setuid/setgid binaries in runtime
  More info: https://pkg.go.dev/vuln/GO-2023-1840
  Standard library
    Found in: runtime@go1.20.4
    Fixed in: runtime@go1.20.5
    Example traces found:
      #1: main.go:12:16: govulncheck.main calls yaml.Unmarshal, which eventually calls runtime.Callers
      #2: main.go:12:16: govulncheck.main calls yaml.Unmarshal, which eventually calls runtime.CallersFrames
      #3: main.go:12:16: govulncheck.main calls yaml.Unmarshal, which eventually calls runtime.Frames.Next
      #4: main.go:13:13: govulncheck.main calls fmt.Println, which eventually calls runtime.GOMAXPROCS
      #5: main.go:4:2: govulncheck.init calls fmt.init, which eventually calls runtime.GOROOT
      #6: main.go:13:13: govulncheck.main calls fmt.Println, which eventually calls runtime.KeepAlive
      #7: main.go:4:2: govulncheck.init calls fmt.init, which eventually calls runtime.SetFinalizer
      #8: main.go:13:13: govulncheck.main calls fmt.Println, which eventually calls runtime.TypeAssertionError.Error
      #9: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.efaceOf
      #10: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.findfunc
      #11: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.float64frombits
      #12: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.forcegchelper
      #13: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.funcMaxSPDelta
      #14: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.lockInit
      #15: main.go:13:13: govulncheck.main calls fmt.Println, which eventually calls runtime.plainError.Error
      #16: main.go:6:2: govulncheck.init calls yaml.init, which eventually calls runtime.throw

Your code is affected by 1 vulnerability from the Go standard library.

Share feedback at https://go.dev/s/govulncheck-feedback.
```

## Run against binary

```sh
go build -o example .
govulncheck -mode=binary example
```

```sh
# same output as in the example above
```