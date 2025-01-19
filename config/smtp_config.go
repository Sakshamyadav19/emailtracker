package config

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
}

// GetSMTPConfig returns the SMTP configuration.
func GetSMTPConfig() SMTPConfig {
	return SMTPConfig{
		Host:     "smtp.gmail.com",       // Replace with your SMTP host
		Port:     "587",                   // Replace with your SMTP port
		Email:    "clickaskipid@gmail.com", // Replace with your email
		Password: "sxrp cjlj kzlo mlgn",    // Replace with your email password
	}
}
