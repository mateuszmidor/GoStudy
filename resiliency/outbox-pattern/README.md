# outbox pattern demo

## Run it

```sh
make rundb # runs postgres with "outbox" table
go run .
```

## psql: run SQL interactively

```sh
make rundb
psql -h localhost -p 5432 -U postgres -d postgres
# enter db password: postgres
# then:
select * from outbox; # semicolon means: EXECUTE NOW
```