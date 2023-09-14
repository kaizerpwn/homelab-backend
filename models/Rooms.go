package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	ID   int
	Name string
}
