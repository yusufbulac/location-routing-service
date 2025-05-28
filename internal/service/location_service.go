package service

import (
	"encoding/json"
	"fmt"
	"github.com/yusufbulac/location-routing-service/internal/cache"
	"github.com/yusufbulac/location-routing-service/internal/logger"
	"github.com/yusufbulac/location-routing-service/internal/model"
	"github.com/yusufbulac/location-routing-service/internal/repository"
	"go.uber.org/zap"
	"math"
	"sort"
	"time"
)

type LocationService interface {
	CreateLocation(location *model.Location) error
	GetAllLocations() ([]model.Location, error)
	GetLocationByID(id uint) (*model.Location, error)
	UpdateLocation(location *model.Location) error
	GetRouteFrom(lat, lng float64) ([]model.Location, error)
	GetPaginatedLocations(limit, offset int) ([]model.Location, error)
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
	err := s.repo.Update(location)
	if err != nil {
		logger.Error("UpdateLocation failed", zap.Error(err), zap.Uint("id", location.ID))
		return err
	}
	return nil
}

func (s *locationService) GetPaginatedLocations(limit, offset int) ([]model.Location, error) {
	return s.repo.GetPaginatedLocations(limit, offset)
}

func (s *locationService) GetRouteFrom(lat, lng float64) ([]model.Location, error) {
	key := fmt.Sprintf("route:%.4f:%.4f", lat, lng)

	// check redis
	if cached, err := cache.Redis.Get(cache.Ctx, key).Result(); err == nil {
		var locations []model.Location
		if err := json.Unmarshal([]byte(cached), &locations); err == nil {
			return locations, nil
		}
	}

	locations, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	sort.Slice(locations, func(i, j int) bool {
		distA := haversine(lat, lng, locations[i].Latitude, locations[i].Longitude)
		distB := haversine(lat, lng, locations[j].Latitude, locations[j].Longitude)
		return distA < distB
	})

	// add cache
	if jsonBytes, err := json.Marshal(locations); err == nil {
		cache.Redis.Set(cache.Ctx, key, jsonBytes, 5*time.Minute)
	}

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
