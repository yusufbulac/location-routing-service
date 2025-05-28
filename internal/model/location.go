package model

import "time"

type Location struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"type:varchar(100);not null" json:"name"`
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
	Color     string  `gorm:"type:char(7);not null" json:"color"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}
