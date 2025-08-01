package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) (bool, error) {
	if err := validate.Struct(s); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return false, err
		}
		return false, errors.New("invalid request data")
	}
	return true, nil
}
