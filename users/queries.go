package users

import (
	"errors"
	"hexnet/api/common"
)

func FindByLogin(login string) (*UserModel, error) {
	var model *UserModel

	common.GetDB().Where("login = ?", login).First(&model)

	if model == nil {
		return nil, errors.New("user not found")
	}

	return model, nil
}

/*
func FindOne(condition interface{}) (*UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return &model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	return db.Save(data).Error
}

func (user *UserModel) Update(data interface{}) error {
	db := common.GetDB()
	return db.Model(user).Updates(data).Error
}

func Create(data interface{}) (*UserModel, error) {
	db := common.GetDB()
	db.Create(data)
	return FindOne(data)
}
*/
