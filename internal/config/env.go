package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Define enums for environment variable keys
const (
	API_URL = "API_URL"
	PORT    = "PORT"
)

// Initialize loads environment variables from .env file
func loadEnv() error {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file, proceeding with system environment variables")
	}

	// Define required environment variables
	requiredEnvVars := []string{
		API_URL,
		PORT,
	}

	var missingEnvVars []string

	// Check for missing environment variables
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			missingEnvVars = append(missingEnvVars, envVar)
		}
	}

	// If there are missing variables, return an error listing them
	if len(missingEnvVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingEnvVars)
	}

	return nil
}

func GetEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
