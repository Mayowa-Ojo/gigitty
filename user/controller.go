package user

import (
	"github.com/gofiber/fiber"
)

// Controller -
type Controller struct {
	UserService IService
}

// NewController -
func NewController(s IService, router fiber.Router) {
	ct := &Controller{
		UserService: s,
	}

	// Register user routes
	router.Get("/", ct.GetAll)
	router.Get("/:id", ct.GetByID)
	router.Post("/", ct.Create)
}

// GetAll -
func (ct *Controller) GetAll(c *fiber.Ctx) {
	users, err := ct.UserService.GetAll(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(users); err != nil {
		c.Status(500).Send(err)
	}
}

// GetByID -
func (ct *Controller) GetByID(c *fiber.Ctx) {
	user, err := ct.UserService.GetByID(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(user); err != nil {
		c.Status(500).Send(err)
	}
}

// Create -
func (ct *Controller) Create(c *fiber.Ctx) {
	user, err := ct.UserService.Create(c)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(user); err != nil {
		c.Status(500).Send(err)
	}
}
