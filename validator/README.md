# validator

- Go library that uses struct tags to specify allowed values
- https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

```go
// Example
type student struct {
	Age    int    `validate:"gte=18,lte=26"`
	Email  string `validate:"required,email"`
	Gender string `validate:"oneof=male female"`
}
```

## Run

```sh
go run .
```

, output:
```text
{Age:30 Email:andrzej#gmail.com Gender:demisexual}
Key: 'student.Age' Error:Field validation for 'Age' failed on the 'lte' tag
Key: 'student.Email' Error:Field validation for 'Email' failed on the 'email' tag
Key: 'student.Gender' Error:Field validation for 'Gender' failed on the 'oneof' tag
```