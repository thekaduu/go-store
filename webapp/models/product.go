package modelsORM

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name        string
	Description string
	Price       float64
	Category    string
	Image       string
}
