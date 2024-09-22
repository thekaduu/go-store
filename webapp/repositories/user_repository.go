package repositories

import (
	"fmt"
	"go-store/application/models"
	"go-store/webapp/config"
	modelsORM "go-store/webapp/models"

	"golang.org/x/crypto/bcrypt"
)

func FindUserById(userId uint) (modelsORM.User, error) {
	var user modelsORM.User

	db := config.ConnectDatabase()
	err := db.First(&user, userId).Error

	return user, err

}

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

func FindByUsername(username string) (modelsORM.User, error) {
	db := config.ConnectDatabase()
	var user modelsORM.User

	err := db.First(&user, "username = ?", username).Error

	return user, err
}

func exists(username string) bool {
	user, err := FindByUsername(username)

	return err == nil && user.ID > 0
}

func SignUp(user models.User) error {
	userModel := modelsORM.User{
		Username: user.Username,
		Password: user.Password,
		Role:     "customer",
	}

	if exists(user.Username) {
		return fmt.Errorf("o usuário \"%s\" já existe", user.Username)
	}

	db := config.ConnectDatabase()

	result := db.Create(&userModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
