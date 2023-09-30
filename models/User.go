package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:64;index" json:"name"`
	Surname  string `gorm:"size:64;index" json:"surname"`
	Email    string `gorm:"size:128;unique;index" json:"email"`
	Image    string `gorm:"size:255" json:"image"`
	Password string `gorm:"size:128;" json:"-"`
	City     string `gorm:"size:128;" json:"city"`
}
