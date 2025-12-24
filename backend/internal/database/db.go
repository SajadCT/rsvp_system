package database

import (
	"log"
	"rsvp/backend/internal/config"
	"rsvp/backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(
		&models.Event{},
		&models.Guest{},
		&models.RSVP{},
	)
	if err != nil {
		log.Fatal("failed to auto migrate database")
	}

	DB = db
}
