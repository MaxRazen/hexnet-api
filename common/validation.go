package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

const (
	ParamLength   = "length"
	ParamFormat   = "format"
	ParamRequired = "required"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Message string     `json:"message"`
	Errors  []ErrorMsg `json:"errors"`
}

func GetValidator() *validator.Validate {
	return binding.Validator.Engine().(*validator.Validate)
}

func IsValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}

func RegisterCustomValidationRules() {
	v := GetValidator()

	err := v.RegisterValidation("login", func(field validator.FieldLevel) bool {
		value := field.Field().String()
		matched, _ := regexp.MatchString("^([-.a-zA-Z\\d]+)$", value)
		return matched
	})

	if err != nil {
		panic("Validation rule `login` could not be registered")
	}
}

func NewValidationError(err error) ValidationError {
	var ve validator.ValidationErrors
	var attr string
	out := make([]ErrorMsg, len(ve))

	if errors.As(err, &ve) {
		for _, fe := range ve {
			attr = fe.Field()
			attr = strings.ToLower(attr[0:1]) + attr[1:]

			out = append(out, ErrorMsg{attr, mapValidationParam(fe.Tag())})
		}
	}
	return ValidationError{
		Message: "Validation error",
		Errors:  out,
	}
}

func NotFoundErrorResponse() map[string]any {
	return gin.H{
		"message": "The record with given ID not found",
	}
}

func mapValidationParam(param string) string {
	switch param {
	case "min":
		return ParamLength
	case "max":
		return ParamLength
	case "login":
		return ParamFormat
	case "required_without":
		return ParamRequired
	}
	return param
}
