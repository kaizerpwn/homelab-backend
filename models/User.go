package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
	City     string
}
