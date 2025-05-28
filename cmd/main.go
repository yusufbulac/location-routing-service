package main

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yusufbulac/location-routing-service/docs"
	"github.com/yusufbulac/location-routing-service/internal/config"
	"github.com/yusufbulac/location-routing-service/internal/handler"
	"github.com/yusufbulac/location-routing-service/internal/middleware"
	"github.com/yusufbulac/location-routing-service/internal/repository"
	"github.com/yusufbulac/location-routing-service/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close DB connection
	sqlDB, err := config.DB.DB()
	if err == nil {
		if cerr := sqlDB.Close(); cerr != nil {
			log.Printf("Error closing DB connection: %v", cerr)
		} else {
			log.Println("Database connection closed")
		}
	}

	log.Println("Server exited")
}
