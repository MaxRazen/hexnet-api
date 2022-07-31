package notes

import "hexnet/api/common"

type NoteCreateData struct {
	Title     string `json:"title" binding:"required_without=Content,max=255"`
	Content   string `json:"content" binding:"required_without=Title,max=65535"`
	CreatedAt int64
	UpdatedAt int64
}

type NoteUpdateData struct {
	Title   string `json:"title" binding:"required_without=Content,max=255"`
	Content string `json:"content" binding:"required_without=Title,max=65535"`
}

func CreateNoteAction(data *NoteCreateData) (*NoteModel, error) {
	model := NoteModel{
		Title:   data.Title,
		Content: data.Content,
	}

	result := common.GetDB().Create(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return &model, nil
}

func UpdateNoteAction(model *NoteModel, data *NoteUpdateData) (*NoteModel, error) {
	modelData := NoteModel{
		Title:   data.Title,
		Content: data.Content,
	}

	result := common.GetDB().Model(model).Updates(modelData)

	return model, result.Error
}

func DeleteNoteAction(id int) error {
	result := common.GetDB().Delete(&NoteModel{}, id)
	return result.Error
}
