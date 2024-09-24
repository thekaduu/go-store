package usersHandler

import (
	"go-store/webapp/handlers/responses"
	"go-store/webapp/repositories"
	"go-store/webapp/requests"
	"go-store/webapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userRequest requests.UserRequest
	var errorResponse responses.ErrorResponse

	err := c.BindJSON(&userRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Occoreu um erro")
	}

	userModel, loginError := repositories.Login(userRequest)

	if loginError != nil {
		errorResponse.Message = loginError.Error()

		utils.NotifyException(loginError, *c.Request, c.ClientIP())

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
