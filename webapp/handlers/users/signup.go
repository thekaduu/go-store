package usersHandler

import (
	"go-store/webapp/repositories"
	"go-store/webapp/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var userRequest requests.UserRequest
	err := c.BindJSON(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ocorreu um erro inesperado.",
		})
		return
	}

	signUpError := repositories.SignUp(userRequest)

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
