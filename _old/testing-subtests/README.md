# Subtests in Go tests

```go
func MultiTest(t *testing.T) {
    sub1func := func(t *testing.T) {
        t.Errorf("this shall not pass!")
    }
    t.Run("subtest1", sub1func)
}
```
