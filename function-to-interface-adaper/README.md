# Function to interface adapter

Such an adapter is a Function that has a member function  
This way we can easily turn regular functions into types compatible with wanted interface  
But that must be single method interface  
Inspired by <https://youtu.be/yeetIgNeIkc?t=502> by Mat Ryer  
Based on http.HandlerFunc:

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
