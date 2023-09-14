package models

import "gorm.io/gorm"

type Devices struct {
	gorm.Model
	ID       int
	Name     string
	Category string
	Room     int
	Status   int8
}
