package utils

import "github.com/go-playground/validator/v10"

func GetErrorMessage(v validator.FieldError) string {
	switch v.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Variable field minimum is " + v.Param()
	case "email":
		return "Format email no valid"
	}
	return "Unknown error"
}
