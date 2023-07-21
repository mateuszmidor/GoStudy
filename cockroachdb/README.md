# CockroachDB

cockroachdb + docker-compose: <https://kb.objectrocket.com/cockroachdb/docker-compose-and-cockroachdb-1151>

## Highlights
- CockroachDB uses PostgreSQL interface, but is actually a key-value DB under the hood
- CockroachDB needs at least 3 nodes to avoid data underreplication
- Web UI at <http://localhost:8080/>

## Play around with cockroach shell (first: ./run_all.sh)

```bash
docker exec -it node_1 /bin/bash
cockroach sql --insecure
```

```sql
\l  -- list databases
create database mydb;
create table users (name varchar(20), age integer);
insert into users (name, age) values ('andrzej', 33);
select * from users;
   name   | age
----------+------
  andrzej |  33
(1 row)

-- the data gets automatically replicated to all 3 nodes
 ```