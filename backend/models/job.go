package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Candidates  []Candidate `json:"candidates" gorm:"foreignKey:JobID"`
}
