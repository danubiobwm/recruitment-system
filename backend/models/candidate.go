package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Resume string `json:"resume"`
	Status string `json:"status"`
	JobID  uint   `json:"jobId"`
}
