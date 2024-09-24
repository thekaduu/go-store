package admin

import (
	models "go-store/webapp/models"
	"go-store/webapp/repositories"
	"go-store/webapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)

	_, err := repositories.CreateProduct(&product)

	if err != nil {
		utils.NotifyException(err, *c.Request, c.ClientIP())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": product})
}
