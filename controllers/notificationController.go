package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/internpulse/renergy-hub-go-backend/pkg"
	"github.com/internpulse/renergy-hub-go-backend/services"
)

// @Summary Get all notifications
// @Tags Notifications
// @Success 200 {array} map[string]interface{}
// @Router /api/v1/notifications [get]
// @Security BearerAuth
func GetNotifications(db *sql.DB, singleUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var notifications interface{}
		var err error

		if singleUser {
			notifications, err = services.GetAllNotificationsForUser(db)
			if err != nil {
				response.Error(c, http.StatusInternalServerError, "failed to get notifications for user")
				return
			}
		} else {
			notifications, err = services.GetAllNotifications(db)
			if err != nil {
				response.Error(c, http.StatusInternalServerError, "failed to get notifications")
				return
			}
		}
		response.Success(c, http.StatusOK, "fetched notifications successfully", notifications)
		return
	}
}

// @Summary Create a new notification
// @Tags Notifications
// @Param       notification body models.Notification true "Book Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/notifications [post]
func CreateNotification(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			UserID  uint   `json:"user_id"`
			Title   string `json:"title"`
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, "an error occured: "+err.Error())
			return
		}

		notification, err := services.CreateNotification(db, req.UserID, req.Title, req.Message)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to create notifications")
			return
		}
		response.Success(c, http.StatusCreated, "created notification successfully", notification)
		return
	}
}

// @Summary Mark notification as read
// @Tags Notifications
// @Param id path uint true "Notification ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/notifications [patch]
func ReadNotification(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		notificationID := uint(0)
		_, err := fmt.Sscanf(id, "%d", &notificationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
			return
		}

		success, err := services.MarkNotificationAsRead(db, notificationID)

		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to mark notifications as read")
		}

		if success {
			c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		}
	}
}

// @Summary Delete a notification
// @Tags Notifications
// @Router /api/v1/notifications [delete]
func DeleteNotification(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		notificationID := uint(0)
		_, err := fmt.Sscanf(id, "%d", &notificationID)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "failed to get id")
		}

		success, err := services.DeleteNotification(db, notificationID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to delete notification")
		}

		if success {
			c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		}
	}
}
