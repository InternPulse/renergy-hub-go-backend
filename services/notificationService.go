package services

import (
	"time"

	"github.com/internpulse/renergy-hub-go-backend/models"
)

var notifications []models.Notification

func GetAllNotifications() []models.Notification {
	return notifications
}

func CreateNotification(userID uint, title, message string) models.Notification {
	notification := models.Notification{
		ID:        uint(len(notifications) + 1),
		UserID:    userID,
		Title:     title,
		Message:   message,
		CreatedAt: time.Now(),
		Read:      false,
	}

	notifications = append(notifications, notification)
	return notification
}

func MarkNotificationAsRead(notificationID uint) bool {
	for i, notification := range notifications {
		if notification.ID == notificationID {
			notifications[i].Read = true
			return true
		}
	}
	return false
}
