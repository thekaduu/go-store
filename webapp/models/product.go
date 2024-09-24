package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	StoreID     uint `gorm:"index"`
	Store       Store
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	CategoryID  uint     `gorm:"index"`
	Category    Category `json:"category"`
	Image       string   `json:"image"`
}
