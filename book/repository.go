package book

import (
	"github.com/Mayowa-Ojo/gigitty/config"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository -
type Repository struct {
	DB *mongo.Database
}

// NewRepository -
func NewRepository(conn config.MongoConn) IRepository {
	return Repository{conn.DB}
}

// GetAll - retrieves all books
func (r Repository) GetAll(ctx *fiber.Ctx) ([]Entity, error) {
	query := bson.D{{}}
	c, err := r.DB.Collection("books").Find(ctx.Fasthttp, query)

	if err != nil {
		ctx.Status(404).Send(err)
	}

	var books = make([]Entity, 0)

	if err := c.All(ctx.Fasthttp, &books); err != nil {
		return books, err
	}

	if err := ctx.JSON(books); err != nil {
		return books, err
	}

	return books, nil
}
