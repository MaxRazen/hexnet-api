package notes

import (
	"hexnet/api/common"
	"log"
)

type NoteModel struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content"`
	CreatedAt int64  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updatedAt" gorm:"column:updated_at"`
}

func AutoMigrate() {
	db := common.GetDB()

	err := db.AutoMigrate(&NoteModel{})

	if err != nil {
		log.Fatal(err.Error())
	}
}
