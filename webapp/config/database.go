package config

import (
	"fmt"
	modelsORM "go-store/webapp/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_DBNAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", user, password, host, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func SetupDatabase() error {
	db := ConnectDatabase()

	return db.AutoMigrate(
		&modelsORM.Store{},
		&modelsORM.User{},
		&modelsORM.Category{},
		&modelsORM.OrderItem{},
		&modelsORM.Order{},
		&modelsORM.Product{},
	)
}
