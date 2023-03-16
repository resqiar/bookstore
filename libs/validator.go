package libs

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func AuthValidator(payload any) string {
	var message string

	// Instantiate validator package
	validate = validator.New()

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				message = err.StructField() + " field is required"
				break
			}
		}
	}

	return message
}
