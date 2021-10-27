package web

import (
	"accounts/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type (
	//AccountHandler deals with http/https requests made to account domain
	AccountHandler struct {
		createAccountUseCase usecase.CreateAccountUseCase
	}

	creationPayload struct {
		DocumentNumber string `json:"documentNumber"`
	}

	creationPresenter struct {
		Id             int64  `json:"id"`
		DocumentNumber string `json:"documentNumber"`
	}
)

//NewAccountHandler builds AccountHandler instance with its dependencies
func NewAccountHandler(createAccountUseCase usecase.CreateAccountUseCase) AccountHandler {
	return AccountHandler{
		createAccountUseCase: createAccountUseCase,
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

	account, err := h.createAccountUseCase.Call(usecase.CreateAccountRequest{DocumentNumber: payload.DocumentNumber})

	if err != nil {
		return err
	}

	presenter := creationPresenter{
		Id:             account.Id,
		DocumentNumber: account.DocumentNumber,
	}

	return c.Status(http.StatusCreated).JSON(presenter)
}

//GetAccount gets an account by a given id query param
func (h AccountHandler) GetAccount(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id param is not a number",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":             id,
		"documentNumber": "1234565",
	})
}
