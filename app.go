package gigitty

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()

	app.Get("/:name", func(c *fiber.Ctx) {
		name := c.Params("name", "John Wick")
		msg := fmt.Sprintf("Hello, %s 👋!", name)
		c.Send(msg)
	})

	app.Listen(4000)
}
