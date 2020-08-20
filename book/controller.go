package book

import (
	"github.com/gofiber/fiber"
)

// Controller -
type Controller struct {
	BookRepository IRepository
}

// NewController -
func NewController(r IRepository, rg fiber.Router) {
	ct := Controller{
		BookRepository: r,
	}

	rg.Get("/", ct.GetAll)
}

// GetAll -
func (ct Controller) GetAll(c *fiber.Ctx) {
	books, err := ct.BookRepository.GetAll(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	c.Status(200).Send(books)
}
