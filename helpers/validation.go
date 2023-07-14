package helpers

import (
	"github.com/go-playground/validator/v10"
)

type ValidationBag struct {
	Field            string `json:"field"`
	FailedValidation string `json:"failed_validation"`
}

var validate = validator.New()

func ValidateStruct(entity interface{}) []*ValidationBag {
	var errors []*ValidationBag

	if err := validate.Struct(entity); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			failedValidation := err.Tag()
			if err.Param() != "" {
				failedValidation += " " + err.Param()
			}
			errors = append(errors,
				&ValidationBag{
					Field:            err.StructNamespace(),
					FailedValidation: failedValidation,
				})
		}
	}
	return errors
}
