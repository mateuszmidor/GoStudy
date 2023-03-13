# mysql driver demo

Based on https://go.dev/doc/tutorial/database-access

## Run

```sh
make run
```

Output:

```text
(...)
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
Album found: {2 Giant Steps John Coltrane 63.99}
(...)
```

## MySQL server

Run MySQL server with admin user `root@admin` and normal user `user@pass` that has access to database `recordings`:

```sh
docker run -d --rm \
    --name=mysql \
    -p=3306:3306 \
    -e=MYSQL_ROOT_PASSWORD=admin \
    -e=MYSQL_USER=user -e=MYSQL_PASSWORD=pass -e=MYSQL_DATABASE=recordings  \
    mysql:8.0.32
```

Insert some data:

```sh
docker exec -it mysql mysqlsh user:pass@localhost:3306 --sql -e "`cat data.sql`"
```

Output:

```text
Cannot set LC_ALL to locale en_US.UTF-8: No such file or directory
WARNING: Using a password on the command line interface can be insecure.

Records: 4  Duplicates: 0  Warnings: 0
```

Retrieve the data:

```sh
docker exec -it mysql mysqlsh user:pass@localhost:3306 --sql -e 'SELECT * from recordings.album'
```

Output:

```text
Cannot set LC_ALL to locale en_US.UTF-8: No such file or directory
WARNING: Using a password on the command line interface can be insecure.
id      title   artist  price
1       Blue Train      John Coltrane   56.99
2       Giant Steps     John Coltrane   63.99
3       Jeru    Gerry Mulligan  17.99
4       Sarah Vaughan   Sarah Vaughan   34.98
```