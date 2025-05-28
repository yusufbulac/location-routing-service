package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbulac/location-routing-service/internal/config"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
