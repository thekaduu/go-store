package repositories

import (
	"go-store/webapp/config"
	modelsORM "go-store/webapp/models"
)

func AllProducts() ([]modelsORM.Product, error) {
	var products []modelsORM.Product
	db := config.ConnectDatabase()

	result := db.Find(&products)

	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}
