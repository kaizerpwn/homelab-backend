package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:64;index"`
	Category string `gorm:"size:64;"`
	RoomID   int
	Room     Room `gorm:"foreignKey:RoomID;"`
	Status   int8
}
