package main

import (
	"log"
	"sql-playground/db"
	"sql-playground/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var (
	engine = html.New("./views", ".html")
)

func main() {
	db.InitDB()

	server := fiber.New(fiber.Config{
		Views: engine,
	})

	server.Static("", "./views/static")

	handlers.InitAPIHandler(server)
	handlers.InitWebHandler(server)

	if err := server.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
