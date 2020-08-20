package book

import (
	"github.com/gofiber/fiber"
)

// Service -
type Service struct {
	BookRepository IRepository
}

// NewService -
func NewService(r IRepository) IService {
	return Service{
		BookRepository: r,
	}
}

// GetAll -
func (ct Service) GetAll(c *fiber.Ctx) ([]Entity, error) {
	books, err := ct.BookRepository.GetAll(c)

	if err != nil {
		return nil, err
	}

	return books, nil
}
