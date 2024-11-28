package services

import (
	"database/sql"
	"time"

	"github.com/internpulse/renergy-hub-go-backend/models"
	_ "github.com/lib/pq"
)

func GetAllNotificationsForUser(db *sql.DB) ([]models.Notification, error) {
	query := `SELECT id, user_id, title, message, created_at, is_read FROM notifications`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Title,
			&notification.Message,
			&notification.CreatedAt,
			&notification.IsRead,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func GetAllNotifications(db *sql.DB) ([]models.Notification, error) {
	query := `SELECT id, user_id, title, message, created_at, is_read FROM notifications`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Title,
			&notification.Message,
			&notification.CreatedAt,
			&notification.IsRead,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

// CreateNotification inserts a new notification into the database.
func CreateNotification(db *sql.DB, userID uint, title, message string) (models.Notification, error) {
	query := `INSERT INTO notifications (user_id, title, message, created_at, is_read) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id uint

	err := db.QueryRow(query, userID, title, message, time.Now(), false).Scan(&id)
	if err != nil {
		return models.Notification{}, err
	}

	return models.Notification{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Message:   message,
		CreatedAt: time.Now(),
		IsRead:    false,
	}, nil
}

// MarkNotificationAsRead marks a notification as read by updating its status in the database.
func MarkNotificationAsRead(db *sql.DB, notificationID uint) (bool, error) {
	query := `UPDATE notifications SET is_read = true WHERE id = $1`
	result, err := db.Exec(query, notificationID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

// DeleteNotification removes a notification from the database.
func DeleteNotification(db *sql.DB, notificationID uint) (bool, error) {
	query := `DELETE FROM notifications WHERE id = $1`
	result, err := db.Exec(query, notificationID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func CreateOrderCreatedNotification(db *sql.DB, userID uint) (models.Notification, error) {
	query := `INSERT INTO notifications (user_id, title, message, created_at, is_read) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id uint

	title := "Order #{orderId} created successfully"
	message := "Your order with id #{orderId} has been created successfully. You can track your order here: "

	err := db.QueryRow(query, userID, title, message, time.Now(), false).Scan(&id)
	if err != nil {
		return models.Notification{}, err
	}

	return models.Notification{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Message:   message,
		CreatedAt: time.Now(),
		IsRead:    false,
	}, nil
}

func CreateOrderShippedNotification(db *sql.DB, userID uint) (models.Notification, error) {
	query := `INSERT INTO notifications (user_id, title, message, created_at, is_read) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id uint

	title := "Order #{orderId} shipped successfully"
	message := "Your order with id #{orderId} has been shipped! You can track your order here: "

	err := db.QueryRow(query, userID, title, message, time.Now(), false).Scan(&id)
	if err != nil {
		return models.Notification{}, err
	}

	return models.Notification{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Message:   message,
		CreatedAt: time.Now(),
		IsRead:    false,
	}, nil
}

func SendEmailVerificationNotification(db *sql.DB, userID uint) (models.Notification, error) {
	query := `INSERT INTO notifications (user_id, title, message, created_at, is_read) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id uint

	title := "Order #{orderId} created successfully"
	message := "Your order with id #{orderId} has been created successfully. You can track your order here: "

	err := db.QueryRow(query, userID, title, message, time.Now(), false).Scan(&id)
	if err != nil {
		return models.Notification{}, err
	}

	return models.Notification{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Message:   message,
		CreatedAt: time.Now(),
		IsRead:    false,
	}, nil
}
