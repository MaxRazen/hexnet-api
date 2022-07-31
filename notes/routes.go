package notes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hexnet/api/common"
	"net/http"
	"strconv"
)

func Routes(router *gin.RouterGroup) {
	router.GET("", getListHandler)
	router.POST("", createHandler)
	router.GET("/:id", getNoteHandler)
	router.PUT("/:id", updateNoteHandler)
	router.DELETE("/:id", deleteNoteHandler)
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

func getListHandler(c *gin.Context) {
	p := common.ExtractPaginationData(c, 0, 20)

	if p.OrderColumn == "" {
		p.OrderColumn = "id"
	}

	if p.OrderDirection == "" {
		p.OrderDirection = "desc"
	}

	records := QueryList(p)
	c.JSON(http.StatusOK, records)
}

func getNoteHandler(c *gin.Context) {
	record, err := extractIdAndFind(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundErrorResponse())
		return
	}

	c.JSON(http.StatusOK, record)
}

func updateNoteHandler(c *gin.Context) {
	record, err := extractIdAndFind(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundErrorResponse())
		return
	}

	var data NoteUpdateData

	if err := c.ShouldBindJSON(&data); err != nil {
		if common.IsValidationError(err) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = UpdateNoteAction(record, &data)

	c.JSON(http.StatusOK, record)
}

func deleteNoteHandler(c *gin.Context) {
	id := extractId(c)

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundErrorResponse())
		return
	}

	err := DeleteNoteAction(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.NotFoundErrorResponse())
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func extractIdAndFind(c *gin.Context) (*NoteModel, error) {
	id := extractId(c)

	if id == 0 {
		return nil, errors.New("invalid id param")
	}
	record, _ := FindById(id)

	if record == nil {
		return nil, errors.New("record not found")
	}

	return record, nil
}

func extractId(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return 0
	}
	return id
}
