package handlers

import (
	"sql-playground/services"

	"github.com/gofiber/fiber/v2"
)

func InitWebHandler(server *fiber.App) {
	data := services.GetAllCities(1)

	server.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Data": data,
		})
	})
}
