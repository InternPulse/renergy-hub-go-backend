// @title Notification Service API
// @version 1.0
// @description API for managing notifications
// @host {host}
// @BasePath /
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/internpulse/renergy-hub-go-backend/config"
	"github.com/internpulse/renergy-hub-go-backend/middlewares"

	"github.com/gin-gonic/gin"

	docs "github.com/internpulse/renergy-hub-go-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	r := gin.Default()
	// routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	if host == "" {
		host = fmt.Sprintf("localhost:%s", port)
	}

	docs.SwaggerInfo.Host = host

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middlewares.LoggerMiddleware())

	r.Run(":" + port)
}
