package main

import (
	"kurbankan/config"
	"kurbankan/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(os.Getenv("APP_PORT"))
}
