# embedded-postgres

- run PostgreSQL server embedded in your application e.g. for testing
- https://github.com/fergusstrange/embedded-postgres

## Default server config

| Field               | Value                                      |
|---------------------|--------------------------------------------|
| Username            | postgres                                   |
| Password            | postgres                                   |
| Database            | postgres                                   |
| Port                | 5432                                       |
| Binaries            | $HOME/.embedded-postgres-go/               |

Yes, the server will download and store dependencies in user's home dir.

## Run

```sh
go run .
```
output:
```sh
Andrzej 32
Jola 24
```

## Run in docker as non-root

```sh
make
```
output:
```sh
uid=1000(regular-user) gid=1000(regular-user) groups=1000(regular-user)
Andrzej 32
Jola 24
```