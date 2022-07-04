package users

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
