package notes

import (
	"github.com/stretchr/testify/assert"
	"hexnet/api/common"
	"testing"
	"time"
)

func setUp(t *testing.T) *assert.Assertions {
	asserts := assert.New(t)
	common.InitTestDbConnection()
	AutoMigrate()
	return asserts
}

func TestCreateNoteAction(t *testing.T) {
	asserts := setUp(t)
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

func TestUpdateNoteAction(t *testing.T) {
	asserts := setUp(t)
	model := createDummyNote()
	data := &NoteUpdateData{Title: "::title::", Content: "::content::"}

	_, err := UpdateNoteAction(model, data)
	asserts.NoError(err, "Error on updating note")
	asserts.Equal(model.Title, data.Title, "Data set was not updated")
	asserts.Greater(model.UpdatedAt, model.CreatedAt, "Updated At must update on every updating")

	data = &NoteUpdateData{Content: "new content"}
	_, err = UpdateNoteAction(model, data)
	asserts.NoError(err, "Updating return an error")
	asserts.Equal(model.Content, data.Content, "Content field was not updated")
}

func TestDeleteNoteAction(t *testing.T) {
	asserts := setUp(t)
	model := createDummyNote()

	err := DeleteNoteAction(int(model.ID))
	asserts.NoError(err, "Deletion should pass without error")

	_, err = FindById(int(model.ID))
	asserts.Error(err, "The record still persists in the DB")
}

func TestFindById(t *testing.T) {
	asserts := setUp(t)
	model := createDummyNote()

	m, err := FindById(int(model.ID))
	asserts.Equal(model, m, "The models must be the same")
	asserts.NoError(err, "Unexpected error has been caught")

	m, err = FindById(999)
	asserts.Nil(m, "Model is persisted but not expected")
	asserts.Error(err, "When record not found an error must be returned")
}

func TestQueryList(t *testing.T) {
	asserts := setUp(t)
	notes := []*NoteModel{
		createDummyNote(),
		createDummyNote(),
		createDummyNote(),
	}

	r := QueryList(common.PaginateData{Limit: 1, Offset: 0, OrderColumn: "id", OrderDirection: "desc"})
	asserts.Len(r, 1)
	asserts.Equal(notes[2], &r[0])

	r = QueryList(common.PaginateData{Limit: 2, Offset: 1, OrderColumn: "id", OrderDirection: "asc"})
	asserts.Len(r, 2)
	asserts.Equal(notes[1], &r[0])
	asserts.Equal(notes[2], &r[1])

	r = QueryList(common.PaginateData{Limit: 20, Offset: 3, OrderColumn: "id", OrderDirection: "asc"})
	asserts.Len(r, 0)
}

func createDummyNote() *NoteModel {
	dt := time.Now().Add(time.Duration(-1) * time.Hour).Unix()
	model := NoteModel{
		Title:     "lorem ipsum",
		CreatedAt: dt,
		UpdatedAt: dt,
	}
	common.GetDB().Create(&model)
	return &model
}
