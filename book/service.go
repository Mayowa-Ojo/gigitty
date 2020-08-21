package book

import (
	"time"

	"github.com/Mayowa-Ojo/gigitty/entity"

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
func (s *Service) GetAll(c *fiber.Ctx) ([]entity.Book, error) {
	books, err := s.BookRepository.GetAll(c)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetByID -
func (s *Service) GetByID(c *fiber.Ctx) (*entity.Book, error) {
	book := new(entity.Book)
	id := c.Params("id")
	bookID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: bookID}}
	book, err = s.BookRepository.GetByID(c, filter)
	if err != nil {
		return nil, err
	}

	return book, nil

}

// Create -
func (s *Service) Create(c *fiber.Ctx) (*entity.Book, error) {
	book := new(entity.Book)

	if err := c.BodyParser(book); err != nil { // parse body into book struct
		return book, err
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

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
