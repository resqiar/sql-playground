package handlers

import (
	"sql-playground/services"

	"github.com/gofiber/fiber/v2"
)

func InitWebHandler(server *fiber.App) {
	data := services.GetAll()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(data)
	})
}
