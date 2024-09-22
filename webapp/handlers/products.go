package productsHandler

import (
	"go-store/webapp/handlers/responses"
	"go-store/webapp/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	products, err := repositories.AllProducts()
	var productsResponse []responses.ProductResponse

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, product := range products {
		productResponse := responses.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Image:       product.Image,
			Price:       product.Price,
		}

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"result":  productsResponse,
	})
}
