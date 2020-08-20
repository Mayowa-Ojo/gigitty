package book

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service -
type Service struct {
	BookRepository IRepository
}

// NewService -
func NewService(r IRepository) IService {
	return &Service{
		BookRepository: r,
	}
}

// GetAll -
func (s *Service) GetAll(c *fiber.Ctx) ([]Entity, error) {
	books, err := s.BookRepository.GetAll(c)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetByID -
func (s *Service) GetByID(c *fiber.Ctx) (*Entity, error) {
	book := new(Entity)
	id := c.Params("id")
	bookID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return book, err
	}

	filter := bson.D{{Key: "_id", Value: bookID}}
	book, err = s.BookRepository.GetByID(c, filter)
	if err != nil {
		return nil, err
	}

	return book, nil

}

// Create -
func (s *Service) Create(c *fiber.Ctx) (*Entity, error) {
	book := new(Entity)

	if err := c.BodyParser(book); err != nil { // parse body into book struct
		return book, err
	}

	result, err := s.BookRepository.Create(c, book)
	if err != nil {
		return book, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	book, err = s.BookRepository.GetByID(c, filter)
	if err != nil {
		return nil, err
	}

	return book, nil
}
