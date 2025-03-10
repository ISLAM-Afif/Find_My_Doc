package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"findMyDoc/internal/entities"
)

func NewPostgresDB(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("Database successfully connected and migrations applied.")
	return db, nil
}

func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Doctor{},
		&entities.Patient{},
		&entities.Appointment{},
	)
	if err != nil {
		return fmt.Errorf("error running migrations: %v", err)
	}
	return nil
}
