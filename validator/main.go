package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// simple validations: https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
type student struct {
	Age    int    `validate:"gte=18,lte=26"`
	Email  string `validate:"required,email"`
	Gender string `validate:"oneof=male female prefer_not_to"`
}

func main() {
	// validator object
	validate := validator.New()

	// struct with all fields wrong :)
	s := student{
		Age:    30,
		Email:  "andrzej#gmail.com",
		Gender: "demisexual",
	}

	// validate the struct
	err := validate.Struct(s)

	// report validation outcome
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", err.(validator.ValidationErrors))
}
