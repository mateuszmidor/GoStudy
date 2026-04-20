# PostgreSQL

PostgreSQL running in docker.
- user: posgres
- pass: pass
- database: postgres
- port: 5432

## Start

```sh
make up # run postgres in docker
make fill # create tables and rows
```

## Stop

```sh
make down
```

## psql - run SQL and terminate

```sh
psql -h localhost -p 5432 -U postgres -d postgres -c 'select * from users'

 id |           name            |             email              |     created_at      |     updated_at
----+---------------------------+--------------------------------+---------------------+---------------------
  1 | Yesenia Schroeder         | domenick20@example.net         | 2025-02-16 22:52:46 | 2025-11-29 05:23:46
  2 | Aurelie Denesik           | jerod72@example.net            | 2025-05-05 04:08:32 | 2025-11-29 05:23:46
  3 | Leonie Schamberger        | kertzmann.dorothea@example.org | 2024-10-15 15:49:13 | 2025-11-29 05:23:46
...
```

## psql - run SQL interactively

```sh
psql -h localhost -p 5432 -U postgres -d postgres
# then:
select * from users; # semicolon means: EXECUTE NOW

```
## SQLs

```sql
select * from pg_tables -- show all tables in "public" schema of the connected database
```

## Useful
- COALESCE(users.name, 'UNKNOWN') - return "UNKNOWN" if users.name is NULL
- WHERE posts.created_at > NOW() - INTERVAL '1 month' - return posts from past 1 month

## Notes
- inner join == join
- left join, right join - they mirror each other
- full outer join - return everything from both tables
