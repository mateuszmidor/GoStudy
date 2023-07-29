# JSON Web Token

Internet-wide authorization bearer token standard.  
https://pkg.go.dev/github.com/go-jose/go-jose/v3@v3.0.0/jwt

## Run

```bash
go run .
```

```json
raw JWT:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiYW5kcnplaiIsImtyeXN0eW5hIl0sImNvbG9yIjoiZ3JlZW4iLCJpc3MiOiJteS1pc3N1ZXIiLCJqdGkiOiIxMjMiLCJuYmYiOjE3MDQwNjcyMDAsInN1YiI6Im15LXN1YmplY3QifQ.3ORzEb3vBWDdJeXrn-h0EDraZI27Q-GOcakjAv6Bmbo

publicClaims:
{Issuer:my-issuer Subject:my-subject Audience:[andrzej krystyna] Expiry:<nil> NotBefore:0xc0000a6470 IssuedAt:<nil> ID:123}

privateClaims:
{Color:green}
```