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

			if err.Tag() == "min" && err.StructField() == "Username" {
				message = err.Field() + " should at least contains 3 characters"
				break
			}

			if err.Tag() == "min" && err.StructField() == "Password" {
				message = err.Field() + " should at least contains 8 characters"
				break
			}
		}
	}

	return message
}
