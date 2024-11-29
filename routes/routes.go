package routes

import (
	"database/sql"

	"github.com/internpulse/renergy-hub-go-backend/controllers"
	"github.com/internpulse/renergy-hub-go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/health", controllers.GetHealth())

	apiv1 := router.Group("/api/v1", middleware.RequiresAuth())
	{
		apiv1.GET("/notifications", controllers.GetNotifications(db, false))
		apiv1.GET("/notifications/:userid", controllers.GetNotifications(db, true))
		apiv1.POST("/notifications", controllers.CreateNotification(db))
		apiv1.PATCH("/notifications/:id/read", controllers.ReadNotification(db))
		apiv1.DELETE("/notifications/:id", controllers.DeleteNotification(db))

		apiv1.POST("/notifications/order-created/:orderId", controllers.OrderCreatedNotification(db))
		apiv1.POST("/notifications/order-shipped/:orderId", controllers.OrderShippedNotification(db))
		apiv1.POST("/notifications/verify-email", controllers.EmailVerificationNotification(db))
	}
}
