package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yusufbulac/location-routing-service/docs"
	"github.com/yusufbulac/location-routing-service/internal/config"
	"github.com/yusufbulac/location-routing-service/internal/handler"
	"github.com/yusufbulac/location-routing-service/internal/middleware"
	"github.com/yusufbulac/location-routing-service/internal/repository"
	"github.com/yusufbulac/location-routing-service/internal/service"
)

// @title Location Routing Service API
// @version 1.0
// @description API for managing and routing locations.
// @host localhost:8080
// @BasePath /
func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// middlewares
	r.Use(middleware.RateLimitMiddleware())

	locationRepo := repository.NewLocationRepository(config.DB)
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/locations", locationHandler.CreateLocation)
	r.GET("/locations", locationHandler.GetAllLocations)
	r.GET("/locations/:id", locationHandler.GetLocationByID)
	r.PUT("/locations/:id", locationHandler.UpdateLocation)

	err := r.Run()
	if err != nil {
		return
	}
}
