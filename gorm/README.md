# gorm

- golang ORM, here backed with embedded-postgres server
- https://medium.com/@itskenzylimon/getting-started-on-golang-gorm-af49381caf3f

## Run

```sh
go run .
```
output:
```sh
2023/12/02 18:28:35 INFO creating user="{ID:0 Name:Andrzej Age:34}"
2023/12/02 18:28:35 INFO created user="{ID:1 Name:Andrzej Age:34}"
2023/12/02 18:28:35 INFO creating user="{ID:0 Name:Jola Age:0}"
2023/12/02 18:28:35 INFO created user="{ID:2 Name:Jola Age:18}"
{ID:1 Name:Andrzej Age:34}
{ID:2 Name:Jola Age:18}
```