package routes

import (
	"fmt"
	"net/http"

	"github.com/internpulse/renergy-hub-go-backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Notification service up and running!"})
	})

	router.GET("/notifications", func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetAllNotifications())
	})

	router.POST("/notifications", func(c *gin.Context) {
		var req struct {
			UserID  uint   `json:"user_id"`
			Title   string `json:"title"`
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		notification := services.CreateNotification(req.UserID, req.Title, req.Message)
		c.JSON(http.StatusCreated, notification)
	})

	router.PUT("/notifications/:id/read", func(c *gin.Context) {
		id := c.Param("id")
		notificationID := uint(0)
		_, err := fmt.Sscanf(id, "%d", &notificationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
			return
		}

		if services.MarkNotificationAsRead(notificationID) {
			c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		}
	})
}
