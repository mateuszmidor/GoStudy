# Developing plugin notes

ID == "" means resource doesnt exist -> terraform remove it from state.  
To implement in read and update funcs:

```go
if resourceDoesntExist {
  d.SetID("")
  return
}
```