package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book - defines the schema of a Book collection
type Book struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title"`
	Author       string             `json:"author"`
	ISBN         string             `json:"isbn"`
	Borrowed     bool               `json:"borrowed"`
	Missing      bool               `json:"missing"`
	BorrowedByID primitive.ObjectID `json:"borrowedById"`
	BorrowedBy   User               `json:"borrowedBy"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}

// User - defines the schema of a User collection
type User struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname     string             `json:"firstname"`
	Lastname      string             `json:"lastname"`
	LibraryID     string             `json:"libraryID"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	BorrowedBooks []Book             `json:"borrowedBooks"`
	// ReturnedBooks
	// FavoriteBooks
}
