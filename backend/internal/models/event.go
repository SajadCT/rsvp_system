package models

import "time"

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	EventDate   time.Time
	CreatedAt   time.Time
}
