package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/me", meHandler)
}

func meHandler(c *gin.Context) {
	jwtData, _ := c.Get("JWT_PAYLOAD")
	c.JSON(http.StatusOK, gin.H{
		"msg":       "hey, it's me",
		"tokenData": jwtData,
	})
}
