package routes

import (
	"github.com/internpulse/renergy-hub-go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", controllers.GetHealth())

	router.GET("/notifications", controllers.GetNotifications())
	router.POST("/notifications", controllers.CreateNotification())
	router.PUT("/notifications/:id/read", controllers.UpdateNotification())
}
