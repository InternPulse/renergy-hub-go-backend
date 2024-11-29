package services

import (
	"database/sql"
	"fmt"

	"github.com/internpulse/renergy-hub-go-backend/models"
	"github.com/lib/pq"
)

func GetUserRole(db *sql.DB) (string, error) {
	return "hi", nil
}
func GetUserDetails(db *sql.DB, userId uint) (models.User, error) {
	query := `SELECT id, name, email FROM dummyusers where id = $1`
	var userDetails models.User

	err := db.QueryRow(query, userId).Scan(&userDetails.ID, &userDetails.FirstName, &userDetails.Email)
	if err == sql.ErrNoRows {
		return models.User{}, fmt.Errorf("user with ID %d not found", userId)
	}
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "42P01" {
				return models.User{}, fmt.Errorf("database table 'dummyusers' does not exist. Please check your database schema")
			}
		}
		return models.User{}, fmt.Errorf("error retrieving user details: %v", err)
	}

	return userDetails, nil
}
