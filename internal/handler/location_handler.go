package handler

import (
	"github.com/yusufbulac/location-routing-service/internal/dto"
	"github.com/yusufbulac/location-routing-service/internal/logger"
	"github.com/yusufbulac/location-routing-service/internal/validation"
	"go.uber.org/zap"
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
		logger.Warn("Invalid JSON received", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid JSON",
		})
		return
	}

	if err := validation.Validator.Struct(req); err != nil {
		logger.Warn("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Validation failed",
			Details: validation.FormatValidationError(err),
		})
		return
	}

	location := model.Location{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Color:     req.Color,
	}

	if err := h.service.CreateLocation(&location); err != nil {
		logger.Error("Could not create location", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Could not create location",
		})
		return
	}

	logger.Info("Location created successfully", zap.String("name", location.Name))
	c.JSON(http.StatusCreated, location)
}

// GetAllLocations godoc
// @Summary List all locations
// @Tags locations
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} model.Location
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/locations [get]
func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, err1 := strconv.Atoi(limitParam)
	offset, err2 := strconv.Atoi(offsetParam)
	if err1 != nil || err2 != nil || limit < 1 || offset < 0 {
		logger.Warn("Invalid pagination parameters", zap.String("limit", limitParam), zap.String("offset", offsetParam))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid pagination parameters",
		})
		return
	}

	locations, err := h.service.GetPaginatedLocations(limit, offset)
	if err != nil {
		logger.Error("Failed to fetch paginated locations", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Could not fetch locations",
		})
		return
	}

	logger.Info("Fetched paginated locations", zap.Int("count", len(locations)))
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
		logger.Warn("Invalid ID parameter", zap.String("id", idParam))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid ID",
		})
		return
	}

	location, err := h.service.GetLocationByID(uint(id))
	if err != nil {
		logger.Warn("Location not found", zap.Int("id", id))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "Location not found",
		})
		return
	}

	logger.Info("Fetched location by ID", zap.Int("id", id))
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
		logger.Warn("Invalid ID parameter", zap.String("id", idParam))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid ID",
		})
		return
	}

	var req dto.LocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid JSON received", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid JSON",
		})
		return
	}

	if err := validation.Validator.Struct(req); err != nil {
		logger.Warn("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Validation failed",
			Details: validation.FormatValidationError(err),
		})
		return
	}

	existing, err := h.service.GetLocationByID(uint(id))
	if err != nil {
		logger.Error("Location not found", zap.Error(err))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "Location not found",
			Details: err.Error(),
		})
		return
	}

	existing.Name = req.Name
	existing.Latitude = req.Latitude
	existing.Longitude = req.Longitude
	existing.Color = req.Color

	if err := h.service.UpdateLocation(existing); err != nil {
		logger.Error("Could not update location", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Could not update location",
		})
		return
	}

	logger.Info("Location updated", zap.Int("id", id))
	c.JSON(http.StatusOK, existing)
}

// GetRoute godoc
// @Summary Get route starting from closest location
// @Tags locations
// @Produce json
// @Param lat query number true "Reference latitude"
// @Param lng query number true "Reference longitude"
// @Success 200 {array} model.Location
// @Failure 400 {object} dto.ErrorResponse
// @Router /api/v1/route [get]
func (h *LocationHandler) GetRoute(c *gin.Context) {
	latParam := c.Query("lat")
	lngParam := c.Query("lng")

	lat, err1 := strconv.ParseFloat(latParam, 64)
	lng, err2 := strconv.ParseFloat(lngParam, 64)

	if err1 != nil || err2 != nil {
		logger.Warn("Invalid lat/lng parameters", zap.String("lat", latParam), zap.String("lng", lngParam))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid lat/lng",
		})
		return
	}

	result, err := h.service.GetRouteFrom(lat, lng)
	if err != nil {
		logger.Error("Failed to compute route", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Could not fetch route",
		})
		return
	}

	logger.Info("Route fetched", zap.Int("count", len(result)))
	c.JSON(http.StatusOK, result)
}
