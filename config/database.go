package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConn -
type MongoConn struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var (
	dbName        = "gigitty"
	mongoURI      = fmt.Sprintf("mongodb://localhost/%s", dbName)
	clientOptions = options.Client().ApplyURI(mongoURI)
)

var conn MongoConn

// Connect - create a mongo connection
func Connect() (MongoConn, error) {
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return conn, err
	}

	conn = MongoConn{
		Client: client,
		DB:     db,
	}

	log.Println("[x] - connected to database")

	return conn, nil
}
