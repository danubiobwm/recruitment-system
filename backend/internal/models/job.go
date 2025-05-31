package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title       string
	Description string
	Status      string
}
