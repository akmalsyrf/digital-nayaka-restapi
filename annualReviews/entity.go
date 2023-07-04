package annualreviews

import (
	"time"

	"gorm.io/gorm"
)

type AnnualReviews struct {
	gorm.Model
	ID         int `gorm:"primary_key"`
	EmployeeID int
	ReviewDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
