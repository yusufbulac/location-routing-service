package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yusufbulac/location-routing-service/internal/model"
)

// --- Mock Repository ---

type mockLocationRepo struct {
	mock.Mock
}

func (m *mockLocationRepo) Create(location *model.Location) error {
	args := m.Called(location)
	return args.Error(0)
}

func (m *mockLocationRepo) FindAll() ([]model.Location, error) {
	args := m.Called()
	return args.Get(0).([]model.Location), args.Error(1)
}

func (m *mockLocationRepo) FindByID(id uint) (*model.Location, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Location), args.Error(1)
}

func (m *mockLocationRepo) Update(location *model.Location) error {
	args := m.Called(location)
	return args.Error(0)
}

// --- Unit Tests ---

func TestCreateLocation(t *testing.T) {
	mockRepo := new(mockLocationRepo)
	service := NewLocationService(mockRepo)

	location := &model.Location{Name: "Test", Latitude: 1.0, Longitude: 1.0, Color: "#FFFFFF"}

	mockRepo.On("Create", location).Return(nil)

	err := service.CreateLocation(location)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllLocations(t *testing.T) {
	mockRepo := new(mockLocationRepo)
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

func TestGetLocationByID_Success(t *testing.T) {
	mockRepo := new(mockLocationRepo)
	service := NewLocationService(mockRepo)

	expected := &model.Location{ID: 1, Name: "Loc", Latitude: 10, Longitude: 20, Color: "#ABCDEF"}

	mockRepo.On("FindByID", uint(1)).Return(expected, nil)

	location, err := service.GetLocationByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, location)
	mockRepo.AssertExpectations(t)
}

func TestGetLocationByID_NotFound(t *testing.T) {
	mockRepo := new(mockLocationRepo)
	service := NewLocationService(mockRepo)

	mockRepo.On("FindByID", uint(999)).Return(&model.Location{}, errors.New("not found"))

	_, err := service.GetLocationByID(999)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateLocation(t *testing.T) {
	mockRepo := new(mockLocationRepo)
	service := NewLocationService(mockRepo)

	location := &model.Location{ID: 1, Name: "Updated", Latitude: 11, Longitude: 22, Color: "#FF00FF"}

	mockRepo.On("Update", location).Return(nil)

	err := service.UpdateLocation(location)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
