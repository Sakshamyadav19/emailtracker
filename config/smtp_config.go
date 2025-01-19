package config

import (
	"log"
	"os"
	
)

type Config struct {
	AuthEmail    string
	AuthPassword string
	SMTPHost     string
	SMTPPort     string
	BaseURL      string
}

func LoadConfig() *Config {
    authEmail := os.Getenv("SMTP_AUTH_EMAIL")
    if authEmail == "" {
        log.Fatal("SMTP_AUTH_EMAIL environment variable not set")
    }
    
    authPassword := os.Getenv("SMTP_AUTH_PASSWORD")
    if authPassword == "" {
        log.Fatal("SMTP_AUTH_PASSWORD environment variable not set")
    }
    
    return &Config{
        AuthEmail:    authEmail,
        AuthPassword: authPassword,
        SMTPHost:     os.Getenv("SMTP_HOST"),
        SMTPPort:     os.Getenv("SMTP_PORT"),
        BaseURL:      os.Getenv("BASE_URL"),
    }
}

