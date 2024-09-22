package middlewares

import (
	"go-store/webapp/handlers/responses"
	"go-store/webapp/repositories"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Message: "Você precisa estar logado para acessar este recurso",
			})

			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("CRYPT_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		user, _ := repositories.FindUserById(userID)

		if !user.IsAdmin() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Você não tem permissão para acessar este recurso"})
			c.Abort()
			return
		}
		c.Set("userId", userID)
		c.Next()
	}
}
