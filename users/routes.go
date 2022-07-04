package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuthRoutes(router *gin.RouterGroup) {
	router.GET("/authorize", routeAuthorize)
}

func routeAuthorize(c *gin.Context) {
	payload := struct {
		Success bool `json:"success"`
	}{Success: true}

	c.JSON(http.StatusOK, payload)
}
