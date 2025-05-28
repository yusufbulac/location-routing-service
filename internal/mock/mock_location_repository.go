package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yusufbulac/location-routing-service/internal/model"
)

// MockLocationRepository is a mocked implementation of the LocationRepository interface.
type MockLocationRepository struct {
	mock.Mock
}

func (m *MockLocationRepository) Create(location *model.Location) error {
	args := m.Called(location)
	return args.Error(0)
}

func (m *MockLocationRepository) FindAll() ([]model.Location, error) {
	args := m.Called()
	return args.Get(0).([]model.Location), args.Error(1)
}

func (m *MockLocationRepository) FindByID(id uint) (*model.Location, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Location), args.Error(1)
}

func (m *MockLocationRepository) Update(location *model.Location) error {
	args := m.Called(location)
	return args.Error(0)
}

func (m *MockLocationRepository) GetPaginatedLocations(limit int, offset int) ([]model.Location, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]model.Location), args.Error(1)
}
