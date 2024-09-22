package usersHandler

import (
	"go-store/application/models"
	"go-store/webapp/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ocorreu um erro inesperado.",
		})
		return
	}

	signUpError := repositories.SignUp(user)

	if signUpError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": signUpError.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário cadastrado com sucesso",
	})
}
