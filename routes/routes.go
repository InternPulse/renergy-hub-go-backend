package routes

import (
	"database/sql"

	"github.com/internpulse/renergy-hub-go-backend/controllers"
	"github.com/internpulse/renergy-hub-go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/health", controllers.GetHealth())

	apiv1Notifications := router.Group("/api/v1/notifications", middleware.RequiresAuth())
	{
		apiv1Notifications.GET("/", controllers.GetNotifications(db, false))
		apiv1Notifications.GET("/:userid", controllers.GetNotifications(db, true))
		apiv1Notifications.POST("/", controllers.CreateNotification(db))
		apiv1Notifications.PATCH("/:id/read", controllers.ReadNotification(db))
		apiv1Notifications.DELETE("/:id", controllers.DeleteNotification(db))

		apiv1Notifications.POST("/order-created/:orderId", controllers.OrderCreatedNotification(db))
		apiv1Notifications.POST("/order-shipped/:orderId", controllers.OrderShippedNotification(db))
		apiv1Notifications.POST("/verify-email", controllers.EmailVerificationNotification(db))
	}

	apiv1Settings := router.Group("/api/v1/settings", middleware.RequiresAuth())
	{
		apiv1Settings.POST("/initialize", controllers.UserSettingsInitialization(db))
		apiv1Settings.GET("/", controllers.GetUsersSettings(db))
		apiv1Settings.PUT("/toggle", controllers.ToggleUserSettings(db))
	}
}
