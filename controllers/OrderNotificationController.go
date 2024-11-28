package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/internpulse/renergy-hub-go-backend/pkg"
	"github.com/internpulse/renergy-hub-go-backend/services"
	"github.com/internpulse/renergy-hub-go-backend/utils"
)

// @Summary creates a database entry for an order creation notification
// @Tags Notifications
// @Success 200 {array} map[string]interface{}
// @Router /api/v1/notifications/order-created [post]
// @Security BearerAuth
func OrderCreatedNotification(db *sql.DB, singleUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserID(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Ensure you are authenticated.")
			return
		}

		notifications, err := services.CreateOrderCreatedNotification(db, userId)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to create notification")

			return
		}
		response.Success(c, http.StatusOK, "Notification sent to user", notifications)

		return
	}
}

// @Summary creates a database entry for an order shipped notification
// @Tags Notifications
// @Success 200 {array} map[string]interface{}
// @Router /api/v1/notifications/order-shipped [post]
// @Security BearerAuth
func OrderShippedNotification(db *sql.DB, singleUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserID(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Ensure you are authenticated.")
			return
		}

		notifications, err := services.CreateOrderShippedNotification(db, userId)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to create notification")
			return
		}
		response.Success(c, http.StatusOK, "Notification sent to user", notifications)
		return
	}
}

// @Summary sends an email verification notification and creates a database entry for it
// @Tags Notifications
// @Success 200 {array} map[string]interface{}
// @Router /api/v1/notifications/verify-email [post]
// @Security BearerAuth
func EmailVerificationNotification(db *sql.DB, singleUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserID(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Ensure you are authenticated.")
			return
		}

		notifications, err := services.SendEmailVerificationNotification(db, userId)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to create notification")
			return
		}
		response.Success(c, http.StatusOK, "Notification sent to user", notifications)
		return
	}
}
