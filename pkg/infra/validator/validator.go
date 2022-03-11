package validator

import (
	"fmt"

	"markets/pkg/app/interfaces"
	valueObjects "markets/pkg/domain/value_objects"

	"github.com/go-playground/validator/v10"
)

type vAlidator struct{}

func (vAlidator) ValidateStruct(m interface{}) []valueObjects.ValidateResult {
	v := validator.New()
	err := v.Struct(m)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	var validatedErros []valueObjects.ValidateResult
	for _, validationErr := range validationErrors {
		message := ""
		if validationErr.Tag() == "required" {
			message = fmt.Sprintf("%s is required", validationErr.Field())
		} else {
			message = fmt.Sprintf("%s invalid %s", validationErr.Field(), validationErr.Tag())
		}
		validatedErros = append(validatedErros, valueObjects.ValidateResult{
			IsValid: false,
			Field:   validationErr.Field(),
			Message: message,
		})
	}

	return validatedErros
}

func NewValidator() interfaces.IValidator {
	return vAlidator{}
}
