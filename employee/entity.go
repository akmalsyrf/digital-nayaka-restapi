package employee

import (
	"time"

	annualreviews "digital-nayaka-test/annualReviews"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID              int `gorm:"primary_key"`
	FirstName       string
	LastName        string
	HireDate        time.Time
	TerminationDate time.Time
	Salary          int
	AnnualReviews   []annualreviews.AnnualReviews
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
