package main

import (
	"accounts/internal/adapter/database/local"
	"accounts/internal/adapter/web"
	"accounts/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()
	accountRepository := local.NewAccountRepository()
	accountHandler := web.NewAccountHandler(usecase.NewCreateAccountUseCase(accountRepository))

	app.Post("/accounts", accountHandler.CreateAccount)
	app.Get("/accounts/:id", accountHandler.GetAccount)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
