package main

import (
	"go-store/webapp/config"
	usersHandler "go-store/webapp/handlers/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setup() {
	godotenv.Load()
	config.SetupDatabase()
}

func setupUserRoutes(app *gin.Engine) {
	app.POST("/login", usersHandler.Login)
	app.POST("/signup", usersHandler.SignUp)
}

func main() {
	setup()

	app := gin.Default()

	setupUserRoutes(app)

	app.Run(":8000")

}
