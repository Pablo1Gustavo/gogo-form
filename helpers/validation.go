package helpers

import (
	"errors"
	"gogo-form/domain"

	"github.com/go-playground/validator/v10"
)

type FailedValidation struct {
	Field            string `json:"field"`
	FailedValidation string `json:"failed_validation"`
}

type ValidationBag []*FailedValidation

var validate = validator.New()

func ValidateStruct(entity interface{}) error {
	var validationErrors ValidationBag

	if err := validate.Struct(entity); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			failedValidation := err.Tag()
			if err.Param() != "" {
				failedValidation += " " + err.Param()
			}

			validationErrors = append(validationErrors,
				&FailedValidation{
					Field:            err.StructNamespace(),
					FailedValidation: failedValidation,
				})
		}
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return &domain.RequestError{
		Code:    422,
		Err:     errors.New("validation error"),
		Details: validationErrors,
	}
}
