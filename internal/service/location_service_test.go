package service

import (
	"errors"
	"github.com/yusufbulac/location-routing-service/internal/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yusufbulac/location-routing-service/internal/model"
)

// --- Unit Tests ---
func TestCreateLocation(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	location := &model.Location{Name: "Test", Latitude: 1.0, Longitude: 1.0, Color: "#FFFFFF"}
	mockRepo.On("Create", location).Return(nil)

	err := service.CreateLocation(location)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllLocations(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	expected := []model.Location{
		{Name: "Loc1", Latitude: 10, Longitude: 20, Color: "#000000"},
		{Name: "Loc2", Latitude: 30, Longitude: 40, Color: "#FFFFFF"},
	}
	mockRepo.On("FindAll").Return(expected, nil)

	locations, err := service.GetAllLocations()
	assert.NoError(t, err)
	assert.Equal(t, expected, locations)
	mockRepo.AssertExpectations(t)
}

func TestGetPaginatedLocations(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	expected := []model.Location{
		{ID: 1, Name: "Pag1", Latitude: 10, Longitude: 10, Color: "#111111"},
		{ID: 2, Name: "Pag2", Latitude: 20, Longitude: 20, Color: "#222222"},
	}
	mockRepo.On("GetPaginatedLocations", 2, 0).Return(expected, nil)

	locations, err := service.GetPaginatedLocations(2, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, locations)
	mockRepo.AssertExpectations(t)
}

func TestGetLocationByID_Success(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	expected := &model.Location{ID: 1, Name: "Loc", Latitude: 10, Longitude: 20, Color: "#ABCDEF"}
	mockRepo.On("FindByID", uint(1)).Return(expected, nil)

	location, err := service.GetLocationByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, location)
	mockRepo.AssertExpectations(t)
}

func TestGetLocationByID_NotFound(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	mockRepo.On("FindByID", uint(999)).Return(&model.Location{}, errors.New("not found"))

	_, err := service.GetLocationByID(999)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateLocation(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	location := &model.Location{ID: 1, Name: "Updated", Latitude: 11, Longitude: 22, Color: "#FF00FF"}
	mockRepo.On("Update", location).Return(nil)

	err := service.UpdateLocation(location)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetRouteFrom(t *testing.T) {
	mockRepo := new(mock.MockLocationRepository)
	service := NewLocationService(mockRepo)

	mockLocations := []model.Location{
		{ID: 1, Name: "A", Latitude: 41.0, Longitude: 28.0},
		{ID: 2, Name: "B", Latitude: 41.1, Longitude: 29.0},
		{ID: 3, Name: "C", Latitude: 41.11, Longitude: 29.01},
	}
	mockRepo.On("FindAll").Return(mockLocations, nil)

	referenceLat := 41.11
	referenceLng := 29.02

	result, err := service.GetRouteFrom(referenceLat, referenceLng)

	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "C", result[0].Name)
	assert.Equal(t, "B", result[1].Name)
	assert.Equal(t, "A", result[2].Name)
	mockRepo.AssertExpectations(t)
}
