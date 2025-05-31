package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto migrate models
	err = db.AutoMigrate(
		&User{},
		&Candidate{},
		&Job{},
		&RecruitmentStrep{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
