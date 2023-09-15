package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:64;index"`
	Category string `gorm:"size:64;"`
	RoomID   int
	Room     Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status   int8
}
