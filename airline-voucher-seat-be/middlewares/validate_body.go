package middlewares

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateBody(input interface{}) map[string]string {
	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		validationErrors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			message := fmt.Sprintf("%s failed on %s", field, e.Tag())
			validationErrors[field] = message
		}

		return validationErrors
	}

	return nil
}
