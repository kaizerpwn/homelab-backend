package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"size:64;index"`
	SquareMeters float64
	HouseID      int
	House        House `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
