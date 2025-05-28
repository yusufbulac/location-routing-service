package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbulac/location-routing-service/internal/config"
	"github.com/yusufbulac/location-routing-service/internal/handler"
	"github.com/yusufbulac/location-routing-service/internal/repository"
	"github.com/yusufbulac/location-routing-service/internal/service"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	locationRepo := repository.NewLocationRepository(config.DB)
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/locations", locationHandler.CreateLocation)
	r.GET("/locations", locationHandler.GetAllLocations)
	r.GET("/locations/:id", locationHandler.GetLocationByID)
	r.PUT("/locations/:id", locationHandler.UpdateLocation)

	err := r.Run()
	if err != nil {
		return
	}
}
