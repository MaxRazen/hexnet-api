package users

import (
	"github.com/gin-gonic/gin"
	"hexnet/api/common"
	"net/http"
)

type AuthorizeRequestData struct {
	Login    string `json:"login" binding:"required,login,min=4"`
	Password string `json:"password" binding:"required,min=6"`
}

func UserAuthRoutes(router *gin.RouterGroup) {
	router.POST("/authorize", routeAuthorize)
}

func routeAuthorize(c *gin.Context) {
	data := &AuthorizeRequestData{}

	if err := c.ShouldBindJSON(&data); err != nil {
		if common.IsValidationError(err) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	payload := struct {
		Success  bool   `json:"success"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}{Success: true, Login: data.Login, Password: data.Password}

	c.JSON(http.StatusOK, payload)
}
