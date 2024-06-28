package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	connectionStr := "user=postgres password=admin dbname=test port=5432 sslmode=disable"
	// Replace with your PostgreSQL connection string details
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
