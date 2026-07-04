# db-migrations

- library for automatic db migration up and down
- https://github.com/golang-migrate/migrate

## Run

```sh
go run .
```

```text
2023/12/08 16:21:50 INFO migration: got schema info version=0 dirty=false
2023/12/08 16:21:50 INFO migration: starting to apply migrations
2023/12/08 16:21:50 INFO migration: migrations have been successfully applied
2023/12/08 16:21:50 INFO migration: got schema info version=2 dirty=false
2023/12/08 16:21:50 INFO closing postgres driver
Andrzej 32 179
Jola 24 164
```

## migration files naming

Must end with `.up.sql` or `.down.sql`