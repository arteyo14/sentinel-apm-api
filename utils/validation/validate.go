package validation

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func FormateValidationError(err error) map[string]string {
	var errors = make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Tag()
	}
	return errors
}

func ValidatePassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
