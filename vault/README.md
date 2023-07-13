# vault

This example shows basic vault client in Go.

After running vault, you can go to localhost:8200 and sign in with token method and `myroot` token.

https://github.com/hashicorp/vault-client-go

## Run

```bash
make # run vault container and run the app
firefox localhost:8200 # login with "myroot" token
make stop # kill vault container
```

```
2023/07/12 21:30:57 secret written successfully
2023/07/12 21:30:57 secret retrieved: map[password1:abc123 password2:correct horse battery staple]
```