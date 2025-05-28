package handler

import (
	"github.com/yusufbulac/location-routing-service/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yusufbulac/location-routing-service/internal/model"
	"github.com/yusufbulac/location-routing-service/internal/service"
)

type LocationHandler struct {
	service service.LocationService
}

func NewLocationHandler(s service.LocationService) *LocationHandler {
	return &LocationHandler{service: s}
}

// POST /locations
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var req dto.LocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatValidationError(err)})
		return
	}

	location := model.Location{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Color:     req.Color,
	}

	if err := h.service.CreateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create location"})
		return
	}

	c.JSON(http.StatusCreated, location)
}

// GET /locations
func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	locations, err := h.service.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

// GET /locations/:id
func (h *LocationHandler) GetLocationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	location, err := h.service.GetLocationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

// PUT /locations/:id
func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req dto.LocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FormatValidationError(err)})
		return
	}

	location := model.Location{
		ID:        uint(id),
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Color:     req.Color,
	}

	if err := h.service.UpdateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update location"})
		return
	}

	c.JSON(http.StatusOK, location)
}
