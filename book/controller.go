package book

import (
	"github.com/Mayowa-Ojo/gigitty/utils"
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
		err := new(fiber.Error)
		err.Code = 404
		err.Message = "[Error]: Resource not found"
		c.Next(err)
		return
	}

	r := utils.NewResponse()
	r.JSONResponse(c, true, 200, "books found", books)
}

// GetByID -
func (ct *Controller) GetByID(c *fiber.Ctx) {
	book, err := ct.BookService.GetByID(c)
	r := utils.NewResponse()

	if err != nil {
		err := new(fiber.Error)
		err.Code = 404
		err.Message = "[Error]: Resource not found"
		c.Next(err)
		return
	}

	r.JSONResponse(c, true, 200, "book found", book)
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
