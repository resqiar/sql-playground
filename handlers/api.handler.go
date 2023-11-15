package handlers

import (
	"sql-playground/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	formatter = message.NewPrinter(language.Indonesian)
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

	api.Get("/get-total-city", func(c *fiber.Ctx) error {
		data, err := services.GetCountCities()
		if err != nil {
			return c.SendString("0")
		}
		return c.SendString(formatter.Sprint(data))
	})

	api.Get("/get-total-population", func(c *fiber.Ctx) error {
		data, err := services.GetSumPopulation()

		if err != nil {
			return c.SendString("0")
		}

		return c.SendString(formatter.Sprint(data))
	})

	api.Get("/get-total-country", func(c *fiber.Ctx) error {
		data, err := services.GetTotalCountry()

		if err != nil {
			return c.SendString("0")
		}

		return c.SendString(formatter.Sprint(data))
	})

	api.Get("/get-total-district", func(c *fiber.Ctx) error {
		data, err := services.GetTotalDistrict()

		if err != nil {
			return c.SendString("0")
		}

		return c.SendString(formatter.Sprint(data))
	})
}
