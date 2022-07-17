package common

import (
	"github.com/stretchr/testify/assert"
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

func TestRegisterCustomValidationRules(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Error("Custom validation rules must bind without panic")
		}
	}()

	RegisterCustomValidationRules()
}

func TestNewValidationError(t *testing.T) {
	asserts := assert.New(t)
	err := makeDummyValidationError()

	asserts.Error(err, "Validation struct does not work")

	errData := NewValidationError(err)
	asserts.Len(errData.Errors, 3, "Validation error must contain an error for each field")
	asserts.Equal("Validation error", errData.Message)
	asserts.Equal("name", errData.Errors[0].Field)
	asserts.Equal(ParamLength, errData.Errors[0].Message)
	asserts.Equal("phoneNumber", errData.Errors[1].Field)
	asserts.Equal("e164", errData.Errors[1].Message)
	asserts.Equal("age", errData.Errors[2].Field)
	asserts.Equal("required", errData.Errors[2].Message)
}

func makeDummyValidationError() error {
	data := struct {
		Name        string `json:"name" binding:"required,min=2,max=32"`
		PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
		Age         int    `json:"age" binding:"required"`
	}{Name: "a", PhoneNumber: "777"}

	return GetValidator().Struct(data)
}
