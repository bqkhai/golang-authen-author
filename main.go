package main

import (
	"authen-author-example/database"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Initialize Database
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	fmt.Println(connectionString)
	database.Connect(connectionString)
	database.Migrate()

	// Initialize Router
	router := initRouter()
	server_port := ":" + os.Getenv("SERVER_PORT")
	router.Run(server_port)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	// api := router.Group("/api")
	return router
}