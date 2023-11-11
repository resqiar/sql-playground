package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := server.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
