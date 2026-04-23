package db

import (
	"multi-tenant-saas-backend-platform/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Organization{},
		&models.User{},
		&models.ProjectMembership{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
