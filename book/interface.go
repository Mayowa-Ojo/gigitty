package book

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
)

// IRepository -
type IRepository interface {
	GetAll(c *fiber.Ctx) ([]Entity, error)
	GetByID(c *fiber.Ctx, filter interface{}) (*Entity, error)
	Create(c *fiber.Ctx, b *Entity) (*mongo.InsertOneResult, error)
}

// IService -
type IService interface {
	GetAll(c *fiber.Ctx) ([]Entity, error)
	GetByID(c *fiber.Ctx) (*Entity, error)
	Create(c *fiber.Ctx) (*Entity, error)
}
