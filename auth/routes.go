package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"hexnet/api/common"
	"hexnet/api/users"
	"log"
	"net/http"
)

type AuthorizeRequestData struct {
	Login    string `json:"login" binding:"required,login,min=4"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthorizeResponseData struct {
	Jwt string `json:"jwt"`
	Exp int64  `json:"exp"`
}

func Routes(router *gin.RouterGroup) {
	router.POST("/authorize", authorizeHandler)
	router.GET("/debug", debugHandler)
}

func authorizeHandler(c *gin.Context) {
	data := &AuthorizeRequestData{}

	if err := c.ShouldBindJSON(&data); err != nil {
		if common.IsValidationError(err) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	model, err := users.FindByLogin(data.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, failResponse())
		return
	}

	if !users.VerifyPassword(data.Password, model.PasswordHash) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, failResponse())
		return
	}

	payload := jwt.MapClaims{
		identityKey: model.ID,
		audienceKey: "users",
	}
	token, expire, err := middleware.TokenGenerator(payload)

	if err != nil {
		log.Println("Auth:authorizeHandler: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, AuthorizeResponseData{
		Jwt: token,
		Exp: expire.Unix(),
	})
}

func debugHandler(c *gin.Context) {
	claims, _ := middleware.GetClaimsFromJWT(c)
	c.JSON(200, gin.H{
		"claims": claims,
		"errors": nil,
	})
}

func failResponse() map[string]any {
	return gin.H{
		"message": "User not found or password mismatch",
	}
}
