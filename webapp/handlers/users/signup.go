package usersHandler

import (
	"go-store/application/models"
	"go-store/webapp/handlers/responses"
	"go-store/webapp/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	errorResponse := responses.ErrorResponse{}

	if err != nil {
		errorResponse.Message = "Ocorreu um erro inesperado."
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	signUpError := repositories.SignUp(user)

	if signUpError != nil {
		errorResponse.Message = "Ocorreu um erro ao cadastrar o usuário."
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := responses.SuccessResponse{Message: "Usuário cadastrado com sucesso"}

	c.JSON(http.StatusCreated, successResponse)
}
