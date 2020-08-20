package main

import (
	"log"

	"github.com/Mayowa-Ojo/gigitty/book"
	"github.com/Mayowa-Ojo/gigitty/config"
	"github.com/gofiber/fiber"
)

func main() {
	conn, err := config.Connect()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	api := app.Group("/api/v1")

	bookRouter := api.Group("/books")
	bookRepository := book.NewRepository(conn)
	book.NewController(bookRepository, bookRouter)

	app.Listen(4000)
}
