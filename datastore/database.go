package datastore

import (
	"database/sql"
	"log"
)

func InitDB(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS notifications (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			message TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			is_read BOOLEAN DEFAULT FALSE,
			user_id INT NOT NULL
		)`

	_, err := db.Exec(query)

	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return err
	}

	log.Println("Database table 'notifications' initialized successfully.")
	return nil
}
