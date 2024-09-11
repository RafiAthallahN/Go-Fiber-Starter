package entities

import "gorm.io/gorm"

type (
	People struct {
		FullName  string `validate:"required"`
		FirstName string
		LastName  string
		Age       int
		Address   string
		Job       string
		IsMarried bool
		gorm.Model
	}
)
