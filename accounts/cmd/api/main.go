package main

import (
	"accounts/internal/adapter/database/dynamo"
	"accounts/internal/adapter/web"
	"accounts/internal/domain/usecase"
	"accounts/internal/infra"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	dynamoClient := infra.ConnectDynamoDB()
	//db, err := infra.ConnectMySQLDB()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer func(db *sql.DB) {
	//	if err := db.Close(); err != nil {
	//		log.Println("Can't close DB Connection")
	//	}
	//}(db)
	app := fiber.New()

	repositoryRegistry := dynamo.NewDynamoDBRepositoryRegistry(dynamoClient) // Replace with "local.NewLocalRepositoryRegistry()" if you want to test local storage

	accountHandler := web.NewAccountHandler(
		usecase.NewCreateAccountUseCase(repositoryRegistry),
		usecase.NewGetAccountUseCase(repositoryRegistry),
	)

	app.Post("/accounts", accountHandler.CreateAccount)
	app.Get("/accounts/:id", accountHandler.GetAccount)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
