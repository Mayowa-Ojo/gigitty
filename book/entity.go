package book

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entity - defines the schema of a Book collection
type Entity struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
	ISBN   int                `json:"isbn"`
	Status string             `json:"status"`
}
