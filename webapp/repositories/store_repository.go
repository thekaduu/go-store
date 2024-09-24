package repositories

import (
	"go-store/webapp/config"
	"go-store/webapp/models"
)

func CreateStore(store *models.Store) (bool, error) {
	db := config.ConnectDatabase()

	result := db.Create(&store)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
