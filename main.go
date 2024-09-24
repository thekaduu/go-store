package main

import (
	"go-store/webapp/config"
	"go-store/webapp/handlers"
	"go-store/webapp/handlers/admin"
	usersHandler "go-store/webapp/handlers/users"
	"go-store/webapp/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func setup() {
	godotenv.Load()
	config.SetupDatabase()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func setupUserRoutes(app *gin.Engine) {
	app.POST("/login", usersHandler.Login)
	app.POST("/signup", usersHandler.SignUp)
}

func setupProductRoutes(app *gin.Engine, routeProtected *gin.RouterGroup) {
	app.GET("/products", handlers.ProductIndex)
	routeProtected.POST("/products", admin.ProductCreate)
}

func setupStoreRoutes(app *gin.Engine) {
	app.POST("/stores", handlers.StoreCreate)
}

func main() {
	setup()

	app := gin.Default()
	protected := app.Group("/admin", middlewares.AuthMiddleware())

	setupUserRoutes(app)
	setupProductRoutes(app, protected)
	setupStoreRoutes(app)

	app.Run(":8000")

}
