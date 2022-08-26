package controller

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	Field string
	Tag   string
	Param string
}

func ValidateRequest(req any) []*ValidationError {
	var errors []*ValidationError
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var ve ValidationError
			ve.Field = err.StructNamespace()
			ve.Tag = err.Tag()
			ve.Param = err.Param()
			errors = append(errors, &ve)
		}
	}
	return errors
}
