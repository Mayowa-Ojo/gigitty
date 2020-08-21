package book

import (
	"github.com/Mayowa-Ojo/gigitty/config"
	"github.com/Mayowa-Ojo/gigitty/entity"
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
	return &Repository{conn.DB}
}

// GetAll - retrieves all books
func (r *Repository) GetAll(ctx *fiber.Ctx) ([]entity.Book, error) {
	var books = make([]entity.Book, 0)
	filter := bson.D{{}}
	c, err := r.DB.Collection("books").Find(ctx.Fasthttp, filter)

	if err != nil {
		return books, err
	}

	if err := c.All(ctx.Fasthttp, &books); err != nil {
		return books, err
	}

	if err := ctx.JSON(books); err != nil {
		return books, err
	}

	return books, nil
}

// GetByID -
func (r *Repository) GetByID(ctx *fiber.Ctx, filter interface{}) (*entity.Book, error) {
	c := r.DB.Collection("books")
	book := new(entity.Book)

	result := c.FindOne(ctx.Fasthttp, filter)

	if err := result.Decode(&book); err != nil {
		return nil, err
	}

	return book, nil
}

// Create -
func (r *Repository) Create(ctx *fiber.Ctx, b *entity.Book) (*mongo.InsertOneResult, error) {
	c := r.DB.Collection("books")

	result, err := c.InsertOne(ctx.Fasthttp, b)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update -
// func (r *Repository) Update(ctx *fiber.Ctx) {

// }

// // Delete -
// func (r *Repository) Delete(ctx *fiber.Ctx) {

// }
