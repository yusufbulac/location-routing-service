package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yusufbulac/location-routing-service/internal/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables only")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("DB_NAME", ""),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&model.Location{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = database
	fmt.Println("Database connection successful")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
