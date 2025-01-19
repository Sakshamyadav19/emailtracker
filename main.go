package main

import (
	"github.com/Sakshamyadav19/emailtracker/handler"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/send", handler.HandleEmailRequest)
	r.GET("/track/:id", handler.HandleTracking)       
	r.GET("/track/count/:id", handler.HandleTrackingCount)

	r.Run(":8080")

}