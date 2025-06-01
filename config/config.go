package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
	AppURL  string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	DBSSL   string
}

// LoadConfig โหลดค่า configuration จากไฟล์ .env
func LoadConfig() *Config {
	// โหลดไฟล์ .env
	err := godotenv.Load()

	// ถ้ามี error ให้ log ออกมา
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
		AppURL:  os.Getenv("APP_URL"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBName:  os.Getenv("DB_NAME"),
		DBSSL:   os.Getenv("DB_SSL"),
	}
}