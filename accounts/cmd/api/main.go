package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type handler struct {
}

func (h handler) Post(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":             10,
		"documentNumber": "1234565",
	})
}

func (h handler) Get(c *fiber.Ctx) error {
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

func main() {

	app := fiber.New()
	h := handler{}

	app.Post("/accounts", h.Post)
	app.Get("/accounts/:id", h.Get)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
