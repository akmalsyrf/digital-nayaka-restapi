package employee

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Employee, error) {
	var employees []Employee

	err := r.db.
		Preload("AnnualReviews").
		Find(&employees).Error

	if err != nil {
		return employees, err
	}

	return employees, nil
}
