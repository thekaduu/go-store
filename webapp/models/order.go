package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	OrderItems           []OrderItem `gorm:"not null"`
	UserID               uint        `gorm:"index"`
	User                 User        `gorm:"not null"`
	StoreID              uint        `gorm:"index"`
	Store                Store       `gorm:"not null"`
	PaymentTransactionId int64       `gorm:"not null"`
	TotalPrice           float64     `gorm:"not null" json:"total_price"`
}
