package handlers

import (
	"fmt"
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
		data, err := services.GetTotal("countrycode")
		if err != nil {
			return c.SendString("0")
		}

		return c.SendString(formatter.Sprint(data))
	})

	api.Get("/get-total-district", func(c *fiber.Ctx) error {
		data, err := services.GetTotal("district")
		if err != nil {
			return c.SendString("0")
		}

		return c.SendString(formatter.Sprint(data))
	})

	api.Get("/filter", func(c *fiber.Ctx) error {
		id := fmt.Sprintf("%%%s%%", c.Query("id", ""))
		name := fmt.Sprintf("%%%s%%", c.Query("name", ""))
		country := fmt.Sprintf("%%%s%%", c.Query("country", ""))
		district := fmt.Sprintf("%%%s%%", c.Query("district", ""))

		data := services.Filter(id, name, country, district)
		return c.JSON(data)
	})

	api.Get("/country-capital", func(c *fiber.Ctx) error {
		rawPage := c.Query("page", "1")

		page, err := strconv.Atoi(rawPage)
		if err != nil {
			return c.JSON(fiber.Map{
				"Error": err,
			})
		}

		data := services.GetAllCountryCapital(page)
		return c.JSON(data)
	})

	api.Get("/filter-country", func(c *fiber.Ctx) error {
		code := fmt.Sprintf("%%%s%%", c.Query("code", ""))
		name := fmt.Sprintf("%%%s%%", c.Query("name", ""))
		capital := fmt.Sprintf("%%%s%%", c.Query("capital", ""))

		data := services.FilterCountry(code, name, capital)
		return c.JSON(data)
	})
}
