package book

import (
	"github.com/Mayowa-Ojo/gigitty/entity"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
)

// IRepository -
type IRepository interface {
	GetAll(c *fiber.Ctx) ([]entity.Book, error)
	GetByID(c *fiber.Ctx, filter interface{}) (*entity.Book, error)
	Create(c *fiber.Ctx, b *entity.Book) (*mongo.InsertOneResult, error)
}

// IService -
type IService interface {
	GetAll(c *fiber.Ctx) ([]entity.Book, error)
	GetByID(c *fiber.Ctx) (*entity.Book, error)
	Create(c *fiber.Ctx) (*entity.Book, error)
}
