package usersHandler

import (
	"go-store/application/models"
	"go-store/webapp/handlers/responses"
	"go-store/webapp/repositories"
	"go-store/webapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var errorResponse responses.ErrorResponse

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Occoreu um erro")
	}

	userModel, loginError := repositories.Login(user)

	if loginError != nil {
		errorResponse.Message = loginError.Error()
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}

	token, tokenErr := utils.GenerateJWT(userModel.ID)

	if tokenErr != nil {
		errorResponse.Message = tokenErr.Error()
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}

	response := responses.SuccessResponse{
		Message: "Ok",
		Result: map[string]any{
			"token": token,
		},
	}

	c.JSON(http.StatusOK, response)
}
