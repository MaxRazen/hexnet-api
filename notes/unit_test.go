package notes

import (
	"github.com/stretchr/testify/assert"
	"hexnet/api/common"
	"testing"
)

func TestCreateNoteAction(t *testing.T) {
	asserts := assert.New(t)
	common.InitTestDbConnection()
	AutoMigrate()

	data := &NoteCreateData{
		Title:   "::title::",
		Content: "::content::",
	}

	model, err := CreateNoteAction(data)
	asserts.NoError(err, "Error on creating a note")
	asserts.NotEmpty(model.ID, "Note ID was not assigned")
	asserts.NotEmpty(model.CreatedAt, "Create At must be filled in the model")
	asserts.Equal(model.CreatedAt, model.UpdatedAt, "Timestamp fields should match after creating")
	asserts.Equal(data.Title, model.Title, "Title field must match with input data")
	asserts.Equal(data.Content, model.Content, "Content field must match with input data")
}
