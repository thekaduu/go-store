package main

import (
	"go-store/webapp/config"
	productsHandler "go-store/webapp/handlers"
	"go-store/webapp/handlers/admin"
	usersHandler "go-store/webapp/handlers/users"
	"go-store/webapp/middlewares"

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

func setupProductRoutes(app *gin.Engine, routeProtected *gin.RouterGroup) {
	app.GET("/products", productsHandler.Index)
	routeProtected.POST("/products", admin.ProductCreate)
}

func main() {
	setup()

	app := gin.Default()
	protected := app.Group("/admin", middlewares.AuthMiddleware())

	setupUserRoutes(app)
	setupProductRoutes(app, protected)

	app.Run(":8000")

}
