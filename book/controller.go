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
	ct := &Controller{
		BookService: s,
	}
	// Register book routes
	router.Get("/", ct.GetAll)
	router.Get("/:id", ct.GetByID)
	router.Post("/", ct.Create)
}

// GetAll -
func (ct *Controller) GetAll(c *fiber.Ctx) {
	books, err := ct.BookService.GetAll(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(books); err != nil {
		c.Status(500).Send(err)
	}
}

// GetByID -
func (ct *Controller) GetByID(c *fiber.Ctx) {
	book, err := ct.BookService.GetByID(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(book); err != nil {
		c.Status(500).Send(err)
	}
}

// Create -
func (ct *Controller) Create(c *fiber.Ctx) {
	book, err := ct.BookService.Create(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(book); err != nil {
		c.Status(500).Send(err)
	}
}
