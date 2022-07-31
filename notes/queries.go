package notes

import (
	"gorm.io/gorm/clause"
	"hexnet/api/common"
)

func FindById(id int) (*NoteModel, error) {
	var model *NoteModel
	result := common.GetDB().First(&model, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func QueryList(p common.PaginateData) []NoteModel {
	var models []NoteModel

	order := clause.OrderByColumn{
		Column: clause.Column{Name: p.OrderColumn},
		Desc:   p.OrderDirection == "desc",
	}

	common.GetDB().Offset(p.Offset).Limit(p.Limit).Order(order).Find(&models)

	return models
}
