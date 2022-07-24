package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexnet/api/auth"
	"hexnet/api/common"
	"hexnet/api/notes"
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
	authMiddleware := auth.NewAuthMiddleware().MiddlewareFunc()

	{ // Users Module
		usersRouteGroup := apiRoutes.Group("/users")
		usersRouteGroup.Use(authMiddleware)
		users.Routes(usersRouteGroup)
	}
	{ // Notes Module
		notesRouteGroup := apiRoutes.Group("/notes")
		notesRouteGroup.Use(authMiddleware)
		notes.Routes(notesRouteGroup)
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
	notes.AutoMigrate()
}

func pingHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"path":   c.Request.RequestURI,
		"result": "PONG!",
	})
}
