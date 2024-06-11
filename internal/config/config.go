package config

import (
	"os"
	"strconv"
)

type ApplicationConfig struct {
	// Server Configuration
	ServerPort string

	// Redis Database Configuration
	CacheHost          string
	CachePort          string
	CachePassword      string
	CacheDatabaseIndex int

	// Email SMTP Configuration
	EmailSMTPServer        string
	EmailSMTPServerPort    int
	EmailSMTPUsername      string
	EmailSMTPPassword      string
	EmailSenderAddress     string
	EmailSenderDisplayName string
}

func LoadConfiguration() *ApplicationConfig {
	// Convert string environment variables to appropriate types
	cacheDatabaseIndex, _ := strconv.Atoi(getEnvironmentVariable("CACHE_DB_INDEX", "0"))
	smtpServerPort, _ := strconv.Atoi(getEnvironmentVariable("EMAIL_SMTP_PORT", "587"))

	return &ApplicationConfig{
		// Server Configuration
		ServerPort: getEnvironmentVariable("SERVER_PORT", "8080"),

		// Redis Cache Configuration
		CacheHost:          getEnvironmentVariable("CACHE_HOST", "localhost"),
		CachePort:          getEnvironmentVariable("CACHE_PORT", "6379"),
		CachePassword:      getEnvironmentVariable("CACHE_PASSWORD", ""),
		CacheDatabaseIndex: cacheDatabaseIndex,

		// Email SMTP Configuration
		EmailSMTPServer:        getEnvironmentVariable("EMAIL_SMTP_SERVER", ""),
		EmailSMTPServerPort:    smtpServerPort,
		EmailSMTPUsername:      getEnvironmentVariable("EMAIL_SMTP_USERNAME", ""),
		EmailSMTPPassword:      getEnvironmentVariable("EMAIL_SMTP_PASSWORD", ""),
		EmailSenderAddress:     getEnvironmentVariable("EMAIL_SENDER_ADDRESS", ""),
		EmailSenderDisplayName: getEnvironmentVariable("EMAIL_SENDER_NAME", "MailFlow Service"),
	}
}

func getEnvironmentVariable(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
