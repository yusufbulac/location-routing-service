package service

import (
	"github.com/yusufbulac/location-routing-service/internal/model"
	"github.com/yusufbulac/location-routing-service/internal/repository"
)

type LocationService interface {
	CreateLocation(location *model.Location) error
	GetAllLocations() ([]model.Location, error)
	GetLocationByID(id uint) (*model.Location, error)
	UpdateLocation(location *model.Location) error
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) CreateLocation(location *model.Location) error {
	return s.repo.Create(location)
}

func (s *locationService) GetAllLocations() ([]model.Location, error) {
	return s.repo.FindAll()
}

func (s *locationService) GetLocationByID(id uint) (*model.Location, error) {
	return s.repo.FindByID(id)
}

func (s *locationService) UpdateLocation(location *model.Location) error {
	return s.repo.Update(location)
}
