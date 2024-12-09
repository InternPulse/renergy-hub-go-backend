package services

import (
	"database/sql"
	"fmt"

	"github.com/internpulse/renergy-hub-go-backend/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func GetSettingsForUser(db *sql.DB, userID uint) (models.Setting, error) {
	query := `
		SELECT	id, 
				user_id, 
				push_notifications, 
				promotions_and_offers, 
				platform_updates, 
				order_updates, 
				transaction_notifications,
				vendor_updates,
				account_and_security_updates
		FROM "Setting"
		WHERE user_id = $1;
		`

	var settings models.Setting

	if err := db.QueryRow(query, userID).Scan(
		&settings.ID,
		&settings.UserID,
		&settings.PushNotifications,
		&settings.PromotionsAndOffers,
		&settings.PlatformUpdates,
		&settings.OrderUpdates,
		&settings.TransactionNotifications,
		&settings.VendorUpdates,
		&settings.AccountAndSecurityUpdates,
	); err != nil {
		return models.Setting{}, err
	}

	return settings, nil
}

func InitializeUserSettings(db *sql.DB, userID uint) (models.Setting, error) {
	query := `
		INSERT INTO "Setting" (
			user_id, 
			push_notifications, 
			promotions_and_offers, 
			platform_updates, 
			order_updates, 
			transaction_notifications, 
			vendor_updates, 
			account_and_security_updates
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`
	var id uint

	err := db.QueryRow(query, userID, true, true, true, true, true, true, true).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return models.Setting{}, fmt.Errorf("you have already initialize your settings %d", userID)
		}
		return models.Setting{}, err
	}

	return models.Setting{
		ID:                        id,
		UserID:                    userID,
		PushNotifications:         true,
		PromotionsAndOffers:       true,
		PlatformUpdates:           true,
		OrderUpdates:              true,
		TransactionNotifications:  true,
		VendorUpdates:             true,
		AccountAndSecurityUpdates: true,
	}, nil
}

func ToggleSetting(db *sql.DB, userID uint, settingsField string) (models.Setting, error) {
	query := fmt.Sprintf(`
		UPDATE "Setting"
		SET %s = NOT COALESCE(%s, false)
		WHERE user_id = $1
		RETURNING id, user_id, push_notifications, promotions_and_offers, platform_updates,
		          order_updates, transaction_notifications, vendor_updates, account_and_security_updates;
	`, settingsField, settingsField)

	var settings models.Setting

	err := db.QueryRow(query, userID).Scan(
		&settings.ID,
		&settings.UserID,
		&settings.PushNotifications,
		&settings.PromotionsAndOffers,
		&settings.PlatformUpdates,
		&settings.OrderUpdates,
		&settings.TransactionNotifications,
		&settings.VendorUpdates,
		&settings.AccountAndSecurityUpdates,
	)
	if err != nil {
		return models.Setting{}, err
	}

	return settings, nil
}
