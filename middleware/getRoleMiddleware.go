package middleware

import (
	"fmt"
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

		user_id, exists := c.Get("user_id")
		fmt.Println(user_id)
		if !exists {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Are you authorized?")
		}

		c.Next()
	}
}
