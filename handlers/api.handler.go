package handlers

import (
	"log"
	"sql-playground/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitAPIHandler(server *fiber.App) {
	api := server.Group("api")

	api.Get("/load-more-city", func(c *fiber.Ctx) error {
		rawPage := c.Query("page", "1")
		page, err := strconv.Atoi(rawPage)
		log.Println(page)
		if err != nil {
			return c.JSON(fiber.Map{
				"Error": err,
			})
		}

		data := services.GetAllCities(page)

		// var builder strings.Builder
		//
		// for _, v := range data {
		// 	fmt.Fprintf(&builder, `
		// 		<tr>
		// 			<td>%d</td>
		// 			<td>%s</td>
		// 			<td>%s</td>
		// 			<td>%s</td>
		// 			<td>%d</td>
		// 		</tr>
		// 	`,
		// 		v.ID, v.Name, v.CountryCode, v.District, v.Population,
		// 	)
		// }

		return c.JSON(data)
	})
}
