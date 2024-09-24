package handlers

import (
	"go-store/webapp/models"
	"go-store/webapp/repositories"
	"go-store/webapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func renderBadRequest(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
}

func StoreCreate(c *gin.Context) {
	var store models.Store
	err := c.ShouldBindJSON(&store)

	if err != nil {
		renderBadRequest(c, "Parâmetros inválidos")
		return
	}

	saved, err := repositories.CreateStore(&store)

	if !saved {
		renderBadRequest(c, "Ocorreu um erro ao tentar criar a loja")
		utils.NotifyException(err, *c.Request, c.ClientIP())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ok"})
}
