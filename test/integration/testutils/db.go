package testutils

import (
	"log"

	"github.com/yusufbulac/location-routing-service/internal/config"
	"github.com/yusufbulac/location-routing-service/internal/model"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func SetupTestDB() {
	TestDB = config.ConnectTestDatabase()
}

func CleanDatabase() {
	if TestDB == nil {
		log.Fatal("TestDB is not initialized. Call SetupTestDB() first.")
	}

	TestDB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	TestDB.Exec("DELETE FROM locations")
	TestDB.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func SeedLocations() {
	if TestDB == nil {
		log.Fatal("TestDB is not initialized. Call SetupTestDB() first.")
	}

	locations := []model.Location{
		{Name: "Point A", Latitude: 40.7128, Longitude: -74.0060, Color: "#ff0000"},
		{Name: "Point B", Latitude: 34.0522, Longitude: -118.2437, Color: "#00ff00"},
		{Name: "Point C", Latitude: 41.8781, Longitude: -87.6298, Color: "#0000ff"},
	}

	for _, loc := range locations {
		if err := TestDB.Create(&loc).Error; err != nil {
			log.Fatalf("failed to seed location: %v", err)
		}
	}
}
