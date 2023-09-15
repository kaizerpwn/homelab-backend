package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:64;index"`
	Surname  string `gorm:"size:64;index"`
	Email    string `gorm:"size:128;unique;index"`
	Password string `gorm:"size:128;"`
	City     string `gorm:"size:128;"`
}
