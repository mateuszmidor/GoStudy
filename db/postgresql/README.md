# PostgreSQL
Based on tutorial: https://databaseschool.com/series/intro-to-postgres/.

PostgreSQL db is running in docker.
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

## psql: run SQL and terminate

```sh
psql -h localhost -p 5432 -U postgres -d postgres -c 'select * from users'
# enter db password: pass

 id |           name            |             email              |     created_at      |     updated_at
----+---------------------------+--------------------------------+---------------------+---------------------
  1 | Yesenia Schroeder         | domenick20@example.net         | 2025-02-16 22:52:46 | 2025-11-29 05:23:46
  2 | Aurelie Denesik           | jerod72@example.net            | 2025-05-05 04:08:32 | 2025-11-29 05:23:46
  3 | Leonie Schamberger        | kertzmann.dorothea@example.org | 2024-10-15 15:49:13 | 2025-11-29 05:23:46
...
```

## psql: run SQL interactively

```sh
psql -h localhost -p 5432 -U postgres -d postgres
# enter db password: pass
# then:
select * from users; # semicolon means: EXECUTE NOW

```
## SQLs

### misc

```sql
-- show all tables in "public" schema of the connected database
select * from pg_tables
```

```sql
-- show all indexes
select * from pg_indexes where schemaname not in ('pg_catalog', 'information_schema');
```

```sql
-- explain what db will do to run the query, eg check if it will use index search or seq scan,
-- use cost to compare qeries, eg with vs without index
explain select * from users where email like '%gmail.com'; -- without actually running the query, outputs less info eg no times
-- OR
explain analyze select * from users where email like '%gmail.com'; -- with actually running the query, outputs more info eg times
```
Output:
```
Seq Scan on users (cost=0.00..1.94 rows=1 width=65)
Filter: ((email)::text ~~ '%gmail.com'::text)
```

```sql
-- returns modified string of "pracownik: Andrzej"
select 'pracownik: ' || name from users
```

```sql
-- create temporary table just for execution of this SQL
with title_content as (select title, content from posts)
select title from title_content;
```

### SQL 101
```sql
-- create table where id will auto-increment
create table cities (id serial primary key, name varchar(100) unique, country varchar(100));
```

```sql
-- insert unique city and return the newly inserted row
insert into cities(name, country) values('katowice', 'polska') returning *;
```

```sql
-- insert or ignore if already exists
insert into cities(name, country) values('łódź', 'polska')
on conflict(name) do nothing;
```

```sql
-- insert or update if already exists -> upsert
insert into cities(name, country) values('katowice', 'Polska')
on conflict(name) do update set country = excluded.Country; -- "excluded" pseudo table holds the row that couldn't get inserted
```

```sql
-- capitalize the city name
update cities set name = initcap(name) where id = 1;
```

```sql
drop table cities;
```

### json, jsonb (binary)
- json - Postgres will store data as is without interpreting
- jsonb - Postgres will parse the data, clean it up (e.g. deduplicate keys, remove irrelevant spaces) and then store it
- `select jsonb_column->'field'` - reference json field and return as jsonb
- `select jsonb_column->>'field'` - reference json field and return as text (always as text)
  - to get number you have to cast it:
      - `select (jsonb_column->>'field)::integer`

```sql
create table computers(id serial primary key, item jsonb default '{}'::jsonb); -- create table with jsonb column
insert into computers(item) values ('{"name":"cpu", "cores":12}'); -- insert first jsonb row
insert into computers(item) values ('{"name":"ram", "size":"48gb"}'); -- insert second jsonb row
select id, item->'name' as item_name from computers; -- read single field from jsonb, return as json
select id, item->>'name' as item_name from computers; -- read single field from jsonb, return as text
select id, (item->>'cores')::integer as cores from computers; -- read single field and cast to integer
update computers set item = item || '{"cores":24}' where item->>'cores' = '12'; -- update jsonb field, "||" here is json merge operator
update computers set item = item - 'cores'; -- update by removing 'cores' jsonb key altogether
```

### auto-incrementing counter using upsert
```sql
create table counters(name varchar(15) unique, val integer)
```

```sql
insert into counters (name, val) values ('views', 1) -- insert counter if doesn't exist yet
on conflict(name) do update set val=counters.val+1 returning *; -- increment counter if exists
```

## Database design

### Types

- **strings**:
  - TEXT - uses only as much space as needed to store the string
  - varchar(100) - to express that some string is by nature limited t oa given lenght
- **integers (always signed)**:
  - SMALLINT - 16 bits signed
  - INT/INTEGER - 32 bits signed
  - BIGINT - 64 bits signed
- **fractionals**:
  - NUMERIC(10, 2) - exact representation, good for money math(10 total digits, including 2 for fractions)
  - DECIMAL(10,2 ) - same as NUMERIC, they are aliases
  - REAL - 32 bits, approximate representation (floating point number, like float32)
  - DOUBLE PRECISION - 64 bits (like float64)
- **serials** - they auto-increment
  - smallserial, serial, bigserial
- **date and time**:
  - date - just date
  - time - just time
  - timestamp - date and time without timezone; BAD
  - timestamptz - timestamp with timezone; postgres will convert and store the provided time as UTC +0
  - timestamp with timezone - same as timestamptz
- **logical***:
  - bool

### Primary keys

Note: Use bigint not to run out of keys

#### DB-generated primary key
Old way:
```sql
create table Companies(
  id bigserial primary key
);
```

New way:
```sql
create table Companies(
  id bigint generated always as identity primary key
);

```
#### Application-generated primary key
Use db type UUIDv7 (time based).


### Constraints

```sql
create table Age(
  value smallint constraint age_must_be_positive check(value >= 0), -- add constraint on row level
  created_at timestamptz default now() ,
  updated_at timestamptz default now() ,
  constraint created_at_less_equal_updated_at check(created_at <= updated_at) -- add constraint on table level; can use OR/AND for complex expressions
);
```

### Foreign key constraint

```sql
create table Companies(
  id bigserial primary key
);

create table Employees(
  id bigserial primary key,
  company_id bigint references Companies(id) -- company_id is foreign key
);
```
Options what to do when the referenced Companies.id is deleted:
- `restrict` - don't allow Companies.id to be deleted (this is default behavior)
- `on delete cascade` - remove all connected Employees
- `on delete set NULL` - make all connected Emplotees orphans (Company-less)
- `on delete set default` - set default Company.id if default is configured for company_id

### Indexes - to avoid full table scans

```sql
-- simple index on a number
create index idx_room on hotel_bookings (room_number);
```

```sql
-- index with support for string pattern matching.
-- WARN: this doesnt work with leading wildcards, eg. email LIKE '%gmail.com' - will use seqential scan
create index idx_email_pattern on users (email text_pattern_ops); -- for VARCHAR columns use: varchar_pattern_ops
explain select * from users where email LIKE 'aubree_%'; -- should use index for search
```

```sql
-- composite index - multicolumn; order makes difference:
-- this index will also cover filtering by fname only, but not by lname only (so called Left-most prefix)
create index idx_fname_lname on users (fname, lname);
```

## Useful
- COALESCE(users.name, 'UNKNOWN') - return "UNKNOWN" if users.name is NULL
- WHERE posts.created_at > NOW() - INTERVAL '1 month' - return posts from past 1 month
- pg_typeof(users.name) - return postgres data type



## Notes
- inner join == join; both values must not be null
- left join - left value can be null
- right join - right value can be null; mirrors left join
- full outer join - return everything from both sides including nulls
- union - removes duplicates from the result
- union all - doesn't remove duplicates


