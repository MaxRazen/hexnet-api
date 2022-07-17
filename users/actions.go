package users

import (
	"hexnet/api/common"
)

type UserCreateData struct {
	Name     string `binding:"required,min=2,max=32"`
	Login    string `binding:"required,login,min=4,max=32"`
	Password string `binding:"required,min=6,max=128"`
}

type UserAuthorizeData struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

type UserAuthorizePayload struct {
	Jwt string `json:"jwt"`
}

func CreateUserAction(data UserCreateData) (*UserModel, error) {
	model := UserModel{
		Name:  data.Name,
		Login: data.Login,
	}

	if err := model.setPassword(data.Password); err != nil {
		return nil, err
	}

	result := common.GetDB().Create(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &model, nil
}
