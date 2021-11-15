package web

import (
	"accounts/internal/domain/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type (
	//AccountHandler deals with http/https requests made to account domain
	AccountHandler struct {
		createAccountUseCase usecase.CreateAccountUseCase
		getAccountUseCase    usecase.GetAccountUseCase
	}

	creationPayload struct {
		DocumentNumber string `json:"document_number"`
	}

	presenter struct {
		Id             string `json:"id"`
		DocumentNumber string `json:"document_number"`
	}
)

//NewAccountHandler builds AccountHandler instance with its dependencies
func NewAccountHandler(createAccountUseCase usecase.CreateAccountUseCase,
	getAccountUseCase usecase.GetAccountUseCase) AccountHandler {
	return AccountHandler{
		createAccountUseCase: createAccountUseCase,
		getAccountUseCase:    getAccountUseCase,
	}
}

//CreateAccount creates an account
func (h AccountHandler) CreateAccount(c *fiber.Ctx) error {
	payload := creationPayload{}

	if err := c.BodyParser(&payload); err != nil {
		log.Println("Error parsing payload:", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	account, err := h.createAccountUseCase.Call(context.Background(), usecase.CreateAccountRequest{DocumentNumber: payload.DocumentNumber})

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	p := presenter{
		Id:             account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return c.Status(http.StatusCreated).JSON(p)
}

//GetAccount gets an account by a given id query param
func (h AccountHandler) GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")

	account, err := h.getAccountUseCase.ById(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	p := presenter{
		Id:             account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return c.Status(fiber.StatusOK).JSON(p)
}
