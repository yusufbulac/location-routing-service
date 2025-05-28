package dto

type LocationRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Color     string  `json:"color" validate:"required,hexcolor"`
}
