package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/internpulse/renergy-hub-go-backend/config"
	response "github.com/internpulse/renergy-hub-go-backend/pkg"
)

func GetRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		_, exists := c.Get("role")
		if !exists {
			response.Error(c, http.StatusUnauthorized, "role not found in request. Are you authorized?")
		}

		c.Next()
	}
}
