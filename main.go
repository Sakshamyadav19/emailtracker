package main

import (
	"log"

	"github.com/Sakshamyadav19/emailtracker/api/config"
	"github.com/Sakshamyadav19/emailtracker/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()
	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {
		handler.HandleEmailRequest(c, cfg)
	})
	r.GET("/track/:id", handler.HandleTracking)
	r.GET("/track-count/:id", handler.HandleTrackingCount)

	r.Run(":8080") // Default listens on :8080
}
