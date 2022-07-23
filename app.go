package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexnet/api/auth"
	"hexnet/api/common"
	"hexnet/api/users"
	"net/http"
)

func prepareApp() common.Config {
	config := common.LoadConfig("")
	common.InitDbConnection(config.Env.DB)
	common.RegisterCustomValidationRules()
	migrate()

	return config
}

func setupServer() (server *gin.Engine) {
	server = gin.Default()

	server.GET("/", pingHandler)
	apiRoutes := server.Group("/api")

	// Auth Module
	auth.Routes(apiRoutes.Group("/auth"))
	authMiddleware := auth.NewAuthMiddleware()

	{ // Users Module
		usersRouteGroup := apiRoutes.Group("/users")
		usersRouteGroup.Use(authMiddleware.MiddlewareFunc())
		users.Routes(usersRouteGroup)
	}

	return server
}

func main() {
	config := prepareApp()

	server := setupServer()
	err := server.Run(config.Env.ServerHost + ":" + config.Env.ServerPort)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Server Finished")
}

func migrate() {
	users.AutoMigrate()
}

func pingHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"path":   c.Request.RequestURI,
		"result": "PONG!",
	})
}
