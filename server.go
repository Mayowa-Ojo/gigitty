package main

import (
	"log"

	"github.com/Mayowa-Ojo/gigitty/book"
	"github.com/Mayowa-Ojo/gigitty/config"
	"github.com/Mayowa-Ojo/gigitty/user"
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
	userRouter := api.Group("/users")

	bookRepository := book.NewRepository(conn)
	userRepository := user.NewRepository(conn)

	bookService := book.NewService(bookRepository, userRepository)
	userService := user.NewService(userRepository)

	book.NewController(bookService, bookRouter)
	user.NewController(userService, userRouter)

	app.Listen(4000)
}
