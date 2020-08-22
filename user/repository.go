package user

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
// func (r *Repository) GetAll(ctx *fiber.Ctx) ([]entity.User, error) {
// 	var users = make([]entity.User, 0)
// 	filter := bson.D{{}}
// 	c, err := r.DB.Collection("users").Find(ctx.Fasthttp, filter)

// 	if err != nil {
// 		return users, err
// 	}

// 	if err := c.All(ctx.Fasthttp, &users); err != nil {
// 		return users, err
// 	}

// 	if err := ctx.JSON(users); err != nil {
// 		return users, err
// 	}

// 	return users, nil
// }

// GetAll - retrieves all books
func (r *Repository) GetAll(ctx *fiber.Ctx) ([]entity.User, error) {
	var users = make([]entity.User, 0)
	// filter := bson.D{{}}
	c := r.DB.Collection("users")

	lookup := bson.D{{"$lookup", bson.D{{"from", "books"}, {"localField", "_id"}, {"foreignField", "borrowedby"}, {"as", "borrowedbooks"}}}}
	// unwind := bson.D{{"$unwind", "$borrowedbooks"}}
	pipeline := mongo.Pipeline{lookup}

	cursor, err := c.Aggregate(ctx.Fasthttp, pipeline)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx.Fasthttp, &users); err != nil {
		return users, err
	}

	if err := ctx.JSON(users); err != nil {
		return users, err
	}

	return users, nil
}

// GetByID -
func (r *Repository) GetByID(ctx *fiber.Ctx, filter interface{}) (*entity.User, error) {
	c := r.DB.Collection("users")
	user := new(entity.User)

	result := c.FindOne(ctx.Fasthttp, filter)

	if err := result.Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

// Create -
func (r *Repository) Create(ctx *fiber.Ctx, u *entity.User) (*mongo.InsertOneResult, error) {
	c := r.DB.Collection("users")

	result, err := c.InsertOne(ctx.Fasthttp, u)

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
