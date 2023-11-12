package main

import (
	"log"
	"sql-playground/db"
	"sql-playground/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDB()

	server := fiber.New()

	handlers.InitWebHandler(server)

	if err := server.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
