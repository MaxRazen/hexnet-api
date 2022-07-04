package common

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetValidator(t *testing.T) {
	inst1 := GetValidator()

	if inst1 == nil {
		t.Error("Validator instance must be returned")
	}

	inst2 := GetValidator()
	if inst1 != inst2 {
		t.Error("GetValidator must be created once")
	}
}

func TestIsValidationError(t *testing.T) {
	err := makeDummyValidationError()

	if !IsValidationError(err) {
		t.Errorf("The error must have validation error type but given %T", err)
	}
}

func TestNewValidatorError(t *testing.T) {
	err := makeDummyValidationError()

	validationError := NewValidatorError(err)

	expected := ValidationError{
		Errors: map[string]string{
			"Age":  "Age is a required field",
			"Name": "Name must be at least 2 characters in length",
		},
	}

	asserts := assert.New(t)
	asserts.Equal(expected, validationError, "Validation errors must be equal be translated translations")
}

func TestGenerateRandomString(t *testing.T) {
	str1 := GenerateRandomString(8, "1234567890")

	if len(str1) != 8 {
		t.Error("Result length unequal to passed argument")
	}

	matched, _ := regexp.MatchString(`^(\d+){8}$`, str1)

	if !matched {
		t.Error("Result contains characters out of the allowed range")
	}

	str2 := GenerateRandomString(100, "")

	if len(str2) != 100 {
		t.Error("Result length unequal to passed argument")
	}

	str3 := GenerateRandomString(0, "")

	if str3 != "" {
		t.Error("Result is not empty")
	}
}

func makeDummyValidationError() error {
	data := struct {
		Name string `validate:"required,min=2,max=32"`
		Age  int    `validate:"required"`
	}{Name: "a"}

	return GetValidator().Struct(data)
}
