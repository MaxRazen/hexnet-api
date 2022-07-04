package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexnet/api/common"
	"hexnet/api/users"
	"net/http"
)

func main() {
	config := common.LoadConfig("")
	common.InitDbConnection(config.Env.DB)
	migrate()

	server := gin.Default()

	server.GET("/", pingHandler)
	apiRoutes := server.Group("/api")
	users.UserAuthRoutes(apiRoutes.Group("/auth"))

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
	payload := struct {
		Path   string `json:"path"`
		Result string `json:"result"`
	}{
		Path:   c.Request.RequestURI,
		Result: "PONG!",
	}

	c.IndentedJSON(http.StatusOK, payload)
}
