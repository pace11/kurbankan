package main

import (
	"kurbankan/config"
	"kurbankan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":4000")
}
