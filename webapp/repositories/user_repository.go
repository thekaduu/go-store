package repositories

import (
	"fmt"
	"go-store/application/models"
	"go-store/webapp/config"
	modelsORM "go-store/webapp/models"

	"golang.org/x/crypto/bcrypt"
)

func Login(userDomain models.User) (modelsORM.User, error) {
	db := config.ConnectDatabase()

	var userModel modelsORM.User

	result := db.Where("username = ?", userDomain.Username).First(&userModel)

	if result.Error != nil {
		return userModel, result.Error
	}

	invalidPassword := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(userDomain.Password))

	if invalidPassword != nil {
		return userModel, fmt.Errorf("usuário ou senha inválido")
	}

	return userModel, nil
}

func SignUp(user models.User) error {
	userModel := modelsORM.User{
		Username: user.Username,
		Password: user.Password,
	}

	db := config.ConnectDatabase()
	result := db.Create(&userModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
