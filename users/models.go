package users

import (
	"errors"
	"hexnet/api/common"
	"hexnet/api/common/hasher"
	"log"
)

const (
	PasswordSaltLength = 32
)

type UserModel struct {
	ID           uint   `gorm:"primary_key"`
	Login        string `gorm:"column:login;unique_index"`
	Name         string `gorm:"column:name"`
	PasswordHash string `gorm:"column:password_hash"`
	CreatedAt    int64  `gorm:"column:created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at"`
}

func AutoMigrate() {
	db := common.GetDB()

	err := db.AutoMigrate(&UserModel{})

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (user *UserModel) setPassword(password string) error {
	if password == "" {
		return errors.New("empty password given")
	}

	salt, err := hasher.GenerateSaltBytes(PasswordSaltLength)

	if err != nil {
		log.Fatal(err)
	}

	hash, err := hasher.Hash(password, salt)

	if err != nil {
		log.Fatal(err)
	}

	user.PasswordHash = string(hash)

	return nil
}

func (user *UserModel) checkPassword(password string) bool {
	matched, err := hasher.VerifyHash(password, user.PasswordHash)
	return matched && err == nil
}
