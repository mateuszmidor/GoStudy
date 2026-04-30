# MCP Joke Server for OpenCode

This simplistic MCP provides programming jokes to OpenCode.

## Run it

```sh
go run .
```

## Add it to OpenCode

```sh
opencode mcp add

# Name: jokes
# Type: Remote
# URL: http://localhost:8080/mcp
```

## Check it

```sh
opencode mcp list

┌  MCP Servers
│
●  ✓ Tell A Joke connected
│      http://localhost:8080/mcp
│
└  1 server(s)
```

## Use it

```sh
opencode run 'tell a joke'
```
anser:

```
> build · minimax-m2.5-free

⚙ Tell_A_Joke_get_joke Unknown

UDP is better in the COVID era since it avoids unnecessary handshakes.
```