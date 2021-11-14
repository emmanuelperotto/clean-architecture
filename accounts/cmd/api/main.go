package main

import (
	"accounts/internal/adapter/database/mysql"
	"accounts/internal/adapter/web"
	"accounts/internal/domain/usecase"
	"accounts/internal/infra"
	"context"
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	_, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("pocs"),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	db, err := infra.ConnectMySQLDB()
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Println("Can't close DB Connection")
		}
	}(db)

	app := fiber.New()
	repositoryRegistry := mysql.NewMySQLRepositoryRegistry(db) // Replace with "local.NewLocalRepositoryRegistry()" if you want to test local storage

	accountHandler := web.NewAccountHandler(
		usecase.NewCreateAccountUseCase(repositoryRegistry),
		usecase.NewGetAccountUseCase(repositoryRegistry),
	)

	app.Post("/accounts", accountHandler.CreateAccount)
	app.Get("/accounts/:id", accountHandler.GetAccount)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
