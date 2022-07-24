package notes

import "hexnet/api/common"

type NoteCreateData struct {
	Title   string `binding:"required,min=2,max=255"`
	Content string `binding:"required,min=1,max=65535"`
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
