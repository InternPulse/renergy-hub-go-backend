package datastore

import (
	"database/sql"
	"log"
)

func InitDB(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS "Notification" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			message TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			is_read BOOLEAN DEFAULT FALSE,
			user_id INT NOT NULL,
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "User"(id) ON DELETE CASCADE
		);
		CREATE TABLE IF NOT EXISTS "Setting" (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			push_notifications BOOLEAN DEFAULT TRUE,
			promotions_and_offers BOOLEAN DEFAULT TRUE,
			platform_updates BOOLEAN DEFAULT TRUE,
			order_updates BOOLEAN DEFAULT TRUE,
			transaction_notifications BOOLEAN DEFAULT TRUE,
			vendor_updates BOOLEAN DEFAULT TRUE,
			account_and_security_updates BOOLEAN DEFAULT TRUE,
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "User"(id) ON DELETE CASCADE,
			CONSTRAINT unique_user UNIQUE (user_id)
		);
	`

	_, err := db.Exec(query)

	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return err
	}

	log.Println("'Notification' and 'Setting' tables initialized successfully.")
	return nil
}
