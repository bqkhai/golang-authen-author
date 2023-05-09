package main

import (
	"authen-author-example/common"
	"authen-author-example/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	common.LoadEnvVariables()
	common.ConnectToDb()
	common.SyncDatabase()
}

func main() {
	// Initialize Router
	router := initRouter()
	server_port := ":" + os.Getenv("SERVER_PORT")
	router.Run(server_port)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	// api := router.Group("/api")
	routes.AuthRoutes(router)
	return router
}
