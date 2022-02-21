package utils

import (
	"github.com/Lenstack/clean-architecture/internal/domain"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(i interface{}) (errors []*domain.Error) {
	var validate = validator.New()
	err := validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			var element domain.Error
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
