package main

import (
	"log"
	"rsvp/backend/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	log.Println("Server running on :8080")
	r.Run(":8080")
}
