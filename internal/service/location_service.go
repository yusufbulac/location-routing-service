package service

import (
	"github.com/yusufbulac/location-routing-service/internal/model"
	"github.com/yusufbulac/location-routing-service/internal/repository"
	"math"
	"sort"
)

type LocationService interface {
	CreateLocation(location *model.Location) error
	GetAllLocations() ([]model.Location, error)
	GetLocationByID(id uint) (*model.Location, error)
	UpdateLocation(location *model.Location) error
	GetRouteFrom(lat, lng float64) ([]model.Location, error)
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

func (s *locationService) GetRouteFrom(lat, lng float64) ([]model.Location, error) {
	locations, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	sort.Slice(locations, func(i, j int) bool {
		distA := haversine(lat, lng, locations[i].Latitude, locations[i].Longitude)
		distB := haversine(lat, lng, locations[j].Latitude, locations[j].Longitude)
		return distA < distB
	})

	return locations, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
