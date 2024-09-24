package repositories

import (
	"go-store/webapp/config"
	models "go-store/webapp/models"
)

func AllProducts() ([]models.Product, error) {
	var products []models.Product
	db := config.ConnectDatabase()

	result := db.Find(&products)

	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func CreateProduct(productModel *models.Product) (bool, error) {
	db := config.ConnectDatabase()

	err := db.Create(productModel).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
