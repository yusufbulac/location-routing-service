package repository

import (
	"github.com/yusufbulac/location-routing-service/internal/model"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(location *model.Location) error
	FindAll() ([]model.Location, error)
	FindByID(id uint) (*model.Location, error)
	Update(location *model.Location) error
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) Create(location *model.Location) error {
	return r.db.Create(location).Error
}

func (r *locationRepository) FindAll() ([]model.Location, error) {
	var locations []model.Location
	err := r.db.Find(&locations).Error
	return locations, err
}

func (r *locationRepository) FindByID(id uint) (*model.Location, error) {
	var location model.Location
	err := r.db.First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *locationRepository) Update(location *model.Location) error {
	return r.db.Save(location).Error
}
