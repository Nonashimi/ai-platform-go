package test

import (
	"project-go/internal/models"

	"gorm.io/gorm"
)

type TestRepository struct {
	db *gorm.DB
}

func NewTestRepo(db *gorm.DB) *TestRepository {
	return &TestRepository{db: db}
}

func (r *TestRepository) CreateTest(test *models.Test) (*models.Test, error) {
	if err := r.db.Create(&test).Error; err != nil {
		return nil, err
	}
	return test, nil
}
