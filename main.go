package main

import (
	"log"
	"os"

	"github.com/internpulse/renergy-hub-go-backend/config"
	"github.com/internpulse/renergy-hub-go-backend/middlewares"
	"github.com/internpulse/renergy-hub-go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	port := os.Getenv("PORT")

	r := gin.Default()

	r.Use(middlewares.LoggerMiddleware())

	routes.RegisterRoutes(r)

	r.Run(":" + port)
}
