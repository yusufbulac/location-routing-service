package main

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yusufbulac/location-routing-service/docs"
	"github.com/yusufbulac/location-routing-service/internal/config"
	"github.com/yusufbulac/location-routing-service/internal/handler"
	"github.com/yusufbulac/location-routing-service/internal/logger"
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
// @BasePath /api/v1
func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	config.ConnectDatabase()

	r := gin.Default()
	r.Use(gin.Recovery())

	// middlewares
	r.Use(middleware.ZapLogger())
	r.Use(middleware.RateLimitMiddleware())

	// dependencies
	locationRepo := repository.NewLocationRepository(config.DB)
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		api.POST("/locations", locationHandler.CreateLocation)
		api.GET("/locations", locationHandler.GetAllLocations)
		api.GET("/locations/:id", locationHandler.GetLocationByID)
		api.PUT("/locations/:id", locationHandler.UpdateLocation)
		api.GET("/route", locationHandler.GetRoute)
	}

	// graceful shutdown setup
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
