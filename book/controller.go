package book

import (
	"github.com/gofiber/fiber"
)

// Controller -
type Controller struct {
	BookService IService
}

// NewController -
func NewController(s IService, router fiber.Router) {
	ct := Controller{
		BookService: s,
	}

	router.Get("/", ct.GetAll)
}

// GetAll -
func (ct Controller) GetAll(c *fiber.Ctx) {
	books, err := ct.BookService.GetAll(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	c.Status(200).Send(books)
}
