package users

import (
	"github.com/go-playground/validator/v10"
	"hexnet/api/common"
)

type UserCreateData struct {
	Name string `validate:"required,min=2,max=32"`
	// TODO: solve the issue with regexp=^([-.a-zA-Z0-9]+)$"
	Login    string `validate:"required,min=4,max=32"`
	Password string `validate:"required,min=6,max=128"`
}

func CreateUserAction(data UserCreateData) (*UserModel, error) {
	var err error

	if err = common.GetValidator().Struct(data); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err.(validator.ValidationErrors)
		}
		return nil, err
	}

	model := UserModel{
		Name:  data.Name,
		Login: data.Login,
	}

	if err = model.setPassword(data.Password); err != nil {
		return nil, err
	}

	result := common.GetDB().Create(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &model, nil
}
