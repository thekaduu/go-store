package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}
