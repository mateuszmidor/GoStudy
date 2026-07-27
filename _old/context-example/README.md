# Context

Inspired by "How to correctly use package context" <https://www.youtube.com/watch?v=-_B5uQ4UGi0>

## Cancellation

Context can be cancelled/timed out, and then it's ```ctx.Done()``` channel gets signalled

## See context types

- WithValue - holds key:value pair
- WithCancel - provides Cancel function
- WithTimeout - automatically timeouts after given period of time and also provides Cancel function
