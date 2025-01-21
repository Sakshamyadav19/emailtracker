package handler

import (
	"net/http"

	"github.com/Sakshamyadav19/emailtracker/config"
	"github.com/Sakshamyadav19/emailtracker/service"
	"github.com/Sakshamyadav19/emailtracker/store"
	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	To      []string `json:"to" binding:"required"`
	Cc      []string `json:"cc"`
	Subject string   `json:"subject" binding:"required"`
	Body    string   `json:"body" binding:"required"`
	IsHTML  bool     `json:"is_html"`
}

var tracker = store.NewTrackerStore()

func HandleEmailRequest(c *gin.Context, cfg *config.Config) {
	var emailReq EmailRequest

	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipients := append(emailReq.To, emailReq.Cc...)
	tracking := service.GenerateTrackingIDs(recipients)

	for _, trackingID := range tracking {
		tracker.AddTrackingID(trackingID)
	}

	emailData := service.EmailData{
		To:       emailReq.To,
		Cc:       emailReq.Cc,
		Subject:  emailReq.Subject,
		Body:     emailReq.Body,
		IsHTML:   emailReq.IsHTML,
		Tracking: tracking,
	}
	
	err := service.SendEmail(cfg, emailData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Email sent successfully",
		"tracking_ids": tracking,
	})
}

func HandleTracking(c *gin.Context) {
	trackingID := c.Param("id")
	if trackingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tracking ID is required"})
		return
	}

	tracker.IncrementOpenCount(trackingID)

	c.Header("Content-Type", "image/gif")
	c.String(http.StatusOK, "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\xff\xff\xff\x21\xf9\x04\x01\x00\x00\x00\x00\x2c\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3b")
}

func HandleTrackingCount(c *gin.Context) {
	trackingID := c.Param("id")
	if trackingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tracking ID is required"})
		return
	}

	count := tracker.GetOpenCount(trackingID)
	c.JSON(http.StatusOK, gin.H{
		"tracking_id": trackingID,
		"open_count":  count,
	})
}
