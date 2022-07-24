package notes

import (
	"github.com/gin-gonic/gin"
	"hexnet/api/common"
	"net/http"
)

//
//type NoteEntity struct {
//	ID uint `json:"id"`
//	Title string `json:"title"`
//	Content  string `json:"content"`
//	CreatedAt uint `json:"createdAt"`
//	UpdatedAt uint `json:"updatedAt"`
//}

func Routes(router *gin.RouterGroup) {
	router.POST("", createHandler)
}

func createHandler(c *gin.Context) {
	data := &NoteCreateData{}

	if err := c.ShouldBindJSON(&data); err != nil {
		if common.IsValidationError(err) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	model, err := CreateNoteAction(data)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Something went wrong during creating a record"},
		)
		return
	}

	c.JSON(http.StatusCreated, model)
}
