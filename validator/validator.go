package validate

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(card interface{}) error {
	err := validate.Struct(card)
	if err != nil {
		return err
	}

	return nil
}
