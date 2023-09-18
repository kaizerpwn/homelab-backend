package models

import "gorm.io/gorm"

type House struct {
	gorm.Model
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Address      string `gorm:"size:128;index"`
	City         string `gorm:"size:128;index"`
	Country      string `gorm:"size:128;index"`
	ZipCode      string `gorm:"size:128;index"`
	Floors       int8
	Garage       bool
	Parking      bool
	SquareMeters float32
	Latitude     float32
	Longitude    float32
	UserID       int
	User         User `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
