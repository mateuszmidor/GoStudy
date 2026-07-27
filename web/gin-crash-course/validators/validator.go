package validators

import (
	"unicode"

	"github.com/go-playground/validator"
)

func ValidateStartsWithCapital(field validator.FieldLevel) bool {
	value := []rune(field.Field().String())
	return len(value) > 0 && unicode.IsUpper(value[0])
}
