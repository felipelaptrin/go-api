package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     int `gorm:"PRIMARY_KEY"`
	Name   string
	Email  string
	Status bool
}
