package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:64;index"`
}
