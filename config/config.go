package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

var AppConfig Config
var DB *gorm.DB

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Could not load .env file, using defaults")
	}

	AppConfig = Config{
		Port:       os.Getenv("PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

func ConnectDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		AppConfig.DbHost, AppConfig.DbUser, AppConfig.DbPassword, AppConfig.DbName, AppConfig.DbPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully!")

	return nil
}
