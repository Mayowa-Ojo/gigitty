package book

import (
	"github.com/gofiber/fiber"
)

// IRepository -
type IRepository interface {
	GetAll(c *fiber.Ctx) ([]Entity, error)
}
