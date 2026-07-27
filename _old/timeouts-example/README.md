# How to handle http client and server side timeouts

- client side - when the server processed the request too long
- server side - when the clien reads back the response too slow

Inspired by "Production-ready Go" <https://youtu.be/YF1qSfkDGAQ?t=1925>

## How it works

- start http server with a handler that waits 3 sec before handling request
- send request with timeout of 1 sec
- timeout happens and handler receives from channel <-req.Context().Done() and abandons request

## Also

Default client and default server have no timeouts and can hang forever!
Here you can see how to create http server with timeouts and http client with timeout.  
