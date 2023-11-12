package handlers

import (
	"sql-playground/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitAPIHandler(server *fiber.App) {
	api := server.Group("api")

	api.Get("/load-more-city", func(c *fiber.Ctx) error {
		rawPage := c.Query("page", "1")

		page, err := strconv.Atoi(rawPage)
		if err != nil {
			return c.JSON(fiber.Map{
				"Error": err,
			})
		}

		data := services.GetAllCities(page)
		return c.JSON(data)
	})
}
