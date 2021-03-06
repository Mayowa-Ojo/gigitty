package book

import (
	"time"

	"github.com/Mayowa-Ojo/gigitty/entity"
	u "github.com/Mayowa-Ojo/gigitty/user"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service -
type Service struct {
	BookRepository IRepository
	UserRepository u.IRepository
}

// NewService -
func NewService(bookRepo IRepository, userRepo u.IRepository) IService {
	return &Service{
		BookRepository: bookRepo,
		UserRepository: userRepo,
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

	if err := book.Validate(); err != nil {
		return nil, err
	}

	id := c.Params("id", "5f4be5f15580c39f9b7a9b4f")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	book.UserID = userID
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
