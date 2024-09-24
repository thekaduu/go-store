package models

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model

	Name         string `json:"name" gorm:"unique"`
	Description  string `json:"description"`
	ProfileImage string `json:"profile_image"`
}
