package validation

import (
	"github.com/go-playground/validator/v10"
)

func FormateValidationError(err error) map[string]string {
	var errors = make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Tag()
	}
	return errors
}
