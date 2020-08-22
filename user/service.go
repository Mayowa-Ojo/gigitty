package user

import (
	"time"

	"github.com/Mayowa-Ojo/gigitty/entity"
	"github.com/Mayowa-Ojo/gigitty/utils"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service -
type Service struct {
	UserRepository IRepository
}

// NewService -
func NewService(r IRepository) IService {
	return &Service{
		UserRepository: r,
	}
}

// GetAll -
func (s *Service) GetAll(c *fiber.Ctx) ([]entity.User, error) {
	users, err := s.UserRepository.GetAll(c)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID -
func (s *Service) GetByID(c *fiber.Ctx) (*entity.User, error) {
	user := new(entity.User)
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: userID}}
	user, err = s.UserRepository.GetByID(c, filter)
	if err != nil {
		return nil, err
	}

	return user, nil

}

// Create -
func (s *Service) Create(c *fiber.Ctx) (*entity.User, error) {
	user := new(entity.User)

	if err := c.BodyParser(user); err != nil { // parse body into user struct
		return nil, err
	}

	libraryID, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}

	user.LibraryID = libraryID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := s.UserRepository.Create(c, user)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	user, err = s.UserRepository.GetByID(c, filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
