package main

import (
	_ "kurbankan/docs"

	"kurbankan/config"
	"kurbankan/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Kurbankan API
// @version         1.0
// @description     API server for Kurbankan qurban management platform.

// @contact.name    Kurbankan Dev
// @contact.email   dev@kurbankan.id

// @host      localhost:4000
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
