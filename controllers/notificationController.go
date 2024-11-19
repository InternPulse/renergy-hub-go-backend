package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/internpulse/renergy-hub-go-backend/services"
)

// @Summary Get all notifications
// @Tags Notifications
// @Success 200 {array} map[string]interface{}
// @Router /notifications [get]
func GetNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "there"})
	}
}

// @Summary Create a new notification
// @Tags Notifications
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /notifications [post]
func CreateNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// @Summary Mark notification as read
// @Tags Notifications
// @Param id path uint true "Notification ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notifications [put]
func UpdateNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}
