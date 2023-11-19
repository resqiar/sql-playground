package handlers

import (
	"sql-playground/services"

	"github.com/gofiber/fiber/v2"
)

func InitWebHandler(server *fiber.App) {
	server.Get("/", func(c *fiber.Ctx) error {
		data := services.GetAllCities(1)
		return c.Render("index", fiber.Map{
			"Data": data,
		})
	})

	server.Get("/country", func(c *fiber.Ctx) error {
		data := services.GetAllCountryCapital(1)
		return c.Render("country", fiber.Map{
			"Data": data,
		})
	})
}
