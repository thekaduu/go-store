package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	Order                 Order   `gorm:"not null"`
	OrderID               uint    `gorm:"index"`
	ProductID             uint    `gorm:"index"`
	Product               Product `gorm:"not null"`
	ProductName           string  `gorm:"not null" json:"product_name"`
	UnitPrice             float64 `gorm:"not null" json:"unit_price"`
	Quantity              int     `gorm:"not null" json:"quantity"`
	TotalPrice            float64 `json:"total_price"`
	UnitPriceWithDiscount float64 `gorm:"default:0" json:"unit_price_with_discount"`
}
