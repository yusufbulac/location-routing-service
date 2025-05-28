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

// CreateLocation godoc
// @Summary Add new location
// @Description Adds a new location with name, coordinates and color
// @Tags locations
// @Accept json
// @Produce json
// @Param location body dto.LocationRequest true "Location JSON"
// @Success 201 {object} model.Location
// @Failure 400 {object} dto.ErrorResponse
// @Router /api/v1/locations [post]
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

// GetAllLocations godoc
// @Summary List all locations
// @Tags locations
// @Produce json
// @Success 200 {array} model.Location
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/locations [get]
func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	locations, err := h.service.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

// GetLocationByID godoc
// @Summary Get location by ID
// @Tags locations
// @Produce json
// @Param id path int true "Location ID"
// @Success 200 {object} model.Location
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/locations/{id} [get]
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

// UpdateLocation godoc
// @Summary Update an existing location
// @Tags locations
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Param location body dto.LocationRequest true "Location JSON"
// @Success 200 {object} model.Location
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/locations/{id} [put]
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
