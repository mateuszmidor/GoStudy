# outbox pattern demo

## Run it

```sh
make rundb
go run .
```

## psql: run SQL interactively

```sh
make rundb
psql -h localhost -p 5432 -U postgres -d postgres
# enter db password: postgress
# then:
select * from outbox; # semicolon means: EXECUTE NOW
```