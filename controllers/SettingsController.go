package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/internpulse/renergy-hub-go-backend/pkg"
	"github.com/internpulse/renergy-hub-go-backend/services"
	"github.com/internpulse/renergy-hub-go-backend/utils"
)

// @Summary Initialize user settings
// @Description Creates default settings entries for a user in the database.
// @Tags Settings
// @Success 200 {object} map[string]interface{} "Default user settings created successfully"
// @Failure 401 {object} map[string]string "Unauthorized. User ID not found in request."
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/settings/initialize [post]
// @Security BearerAuth
func UserSettingsInitialization(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserID(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Ensure you are authenticated.")
			return
		}

		settings, err := services.InitializeUserSettings(db, userId)
		if err != nil {
			response.Error(c, http.StatusBadRequest, fmt.Sprintf("failed to initialize user's settings: %v", err))
			return
		}
		response.Success(c, http.StatusOK, "User's settings initialized successfully", settings)

		return
	}
}

// @Summary Retrieve user settings
// @Description Fetches the saved settings for the authenticated user.
// @Tags Settings
// @Success 200 {object} map[string]interface{} "Fetched user settings successfully"
// @Failure 401 {object} map[string]string "Unauthorized. User ID not found in request."
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/settings [get]
// @Security BearerAuth
func GetUsersSettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserID(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Ensure you are authenticated.")
			return
		}

		settings, err := services.GetSettingsForUser(db, userId)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, fmt.Sprintf("failed to get user's settings: %v", err))
			return
		}
		response.Success(c, http.StatusOK, "User's settings fetched successfully", settings)

		return
	}
}

// @Summary Toggle one (only one) specific user setting
// @Description Toggles a setting based on a query parameter. The query parameter specifies which setting to toggle. You can toggle only one at a time. Note that the query parameter doesn't need to be provided a valid, since it is a toggle, just the key.
// @Tags Settings
// @Param push_notifications query bool false "Toggle push notifications setting"
// @Param promotions_and_offers query bool false "Toggle promotions and offers setting"
// @Param platform_updates query bool false "Toggle platform updates setting"
// @Param order_updates query bool false "Toggle order updates setting"
// @Param transaction_notifications query bool false "Toggle transaction notifications setting"
// @Param vendor_updates query bool false "Toggle vendor updates setting"
// @Param account_and_security_updates query bool false "Toggle account and security updates setting"
// @Success 200 {object} map[string]interface{} "Toggled user setting successfully"
// @Failure 400 {object} map[string]string "Invalid query parameter for toggling settings"
// @Failure 401 {object} map[string]string "Unauthorized. User ID not found in request."
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/settings/toggle [put]
// @Security BearerAuth
func ToggleUserSettings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserID(c)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "User id not found in request. Ensure you are authenticated.")
			return
		}

		fields := [7]string{
			"push_notifications",
			"promotions_and_offers",
			"platform_updates",
			"order_updates",
			"transaction_notifications",
			"vendor_updates",
			"account_and_security_updates",
		}

		queryParams := c.Request.URL.Query()
		firstParam := make(map[string]string)

		var firstKey string
		for key, values := range queryParams {
			if len(values) > 0 {
				firstParam[key] = values[0]
				firstKey = key
				break
			}
		}

		isValid := false
		for _, field := range fields {
			if field == firstKey {
				isValid = true
				break
			}
		}

		if !isValid {
			response.Error(c, http.StatusBadRequest, fmt.Sprintf("Query parameter %s not a valid key for toggling settings", firstKey))
			return
		}

		settings, err := services.ToggleSetting(db, userId, firstKey)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, fmt.Sprintf("failed to toggle user's settings: %v", err))
			return
		}
		response.Success(c, http.StatusOK, "Successfully toggled user's settings", settings)

		return
	}
}
