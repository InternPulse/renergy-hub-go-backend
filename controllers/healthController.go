package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health Check
// @Tags Health
// @Success 200 {object} map[string]interface{}
// @Router      /health [get]
func GetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Notification service up and running!"})
	}
}
