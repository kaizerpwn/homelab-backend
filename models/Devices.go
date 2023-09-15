package models

import "gorm.io/gorm"

type Devices struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:64;index"`
	Category string `gorm:"size:64;"`
	Room     int
	Status   int8
}
