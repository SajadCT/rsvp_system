package main

import (
	"log"
	"rsvp/backend/internal/database"
)

func main() {
	database.ConnectDB()
	log.Println("Database connected successfully")
}
