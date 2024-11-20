package routes

import (
	"database/sql"

	"github.com/internpulse/renergy-hub-go-backend/controllers"
	"github.com/internpulse/renergy-hub-go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/health", controllers.GetHealth())

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/notifications", middleware.RequiresAuth(), controllers.GetNotifications(db, false))
		apiv1.GET("/notifications/:userid", middleware.RequiresAuth(), controllers.GetNotifications(db, true))
		apiv1.POST("/notifications", middleware.RequiresAuth(), controllers.CreateNotification(db))
		apiv1.PATCH("/notifications/:id/read", middleware.RequiresAuth(), controllers.ReadNotification(db))
		apiv1.DELETE("/notifications/:id", middleware.RequiresAuth(), controllers.DeleteNotification(db))
	}
}
