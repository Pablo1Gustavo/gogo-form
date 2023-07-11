package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
const timeoutTime = 10 * time.Second

func InitDB() {
	connectionURI := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		"root",
		"example",
		"127.0.0.1",
		"27017")
		
	clientOpts := options.Client().ApplyURI(connectionURI)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutTime)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("gogo-form")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
