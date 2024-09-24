package modelsORM

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Profile     string `json:"profile_pic"`
}
