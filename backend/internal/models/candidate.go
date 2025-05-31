package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name   string
	Email  string
	Phone  string
	Resume string
	Status string
	JobID  uint
	Job    Job `gorm:"foreignKey:JobID"`
}
