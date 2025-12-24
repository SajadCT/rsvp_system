package models

import "time"

type RSVP struct {
	ID        uint   `gorm:"primaryKey"`
	EventID   uint   `gorm:"not null"`
	GuestID   uint   `gorm:"not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
}
