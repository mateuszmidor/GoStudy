# GO "errors" package study

GO blog: <https://blog.golang.org/go1.13-errors>

```go
- fmt.Errorf("%w", wrappedError) // wrap wrappedError into new error with description
- errors.New(description) // create new error with description
- errors.Is(errorChain, chainElement) // does errorChain contain chainElement?
- errors.As(errorChain, dstErrorPtr) // try extract dstErrorPtr type from errorChain
- errors.Unwrap(errorChain) // take off outmost layer of the chain
```