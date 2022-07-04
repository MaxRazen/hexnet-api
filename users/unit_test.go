package users

import (
	"github.com/stretchr/testify/assert"
	"hexnet/api/common"
	"testing"
)

func newUserModel() UserModel {
	return UserModel{
		ID:    2,
		Login: "john.doe",
		Name:  "John Doe",
	}
}

func TestCreateUserAction(t *testing.T) {
	asserts := assert.New(t)
	common.InitTestDbConnection()
	AutoMigrate()

	data := UserCreateData{
		Name:     "John Dou",
		Login:    "j.dou",
		Password: "pa$$w0rd",
	}

	model, err := CreateUserAction(data)

	asserts.NotEmpty(model.ID, "User ID was not assigned")
	asserts.NoError(err, "Error due creating user")

	data.Name = "1"
	data.Login = "$"
	data.Password = ""
	_, err = CreateUserAction(data)
	asserts.Error(err, "User created with invalid data")
	asserts.True(common.IsValidationError(err), "The error type must be validation error")
}

func TestUserModel(t *testing.T) {
	asserts := assert.New(t)

	// Testing UserModel's password feature
	userModel := newUserModel()
	err := userModel.setPassword("pa$w0rd")

	asserts.NoError(err, "Setting password should not return an error")
	asserts.NotEmpty(userModel.PasswordHash, "Empty password hash should not be empty")
	asserts.True(userModel.checkPassword("pa$w0rd"), "Password must match with hash")
	asserts.False(userModel.checkPassword("password"), "Match of different password must return false")

	err = userModel.setPassword("")
	asserts.Error(err, "Setting empty password must return an error")
}
