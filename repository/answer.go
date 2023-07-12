package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"gogo-form/database"
	"gogo-form/models"
)

type AnswerRepository struct {
	collection *mongo.Collection
}

func NewAnswerRepository() *AnswerRepository {
	collection := database.GetCollection("answers")
	return &AnswerRepository{collection}
}

func (repo *AnswerRepository) Create(ctx context.Context, answer models.Answer) (*mongo.InsertOneResult, error) {
	return repo.collection.InsertOne(ctx, answer)
}

func (repo *AnswerRepository) GetAll(ctx context.Context) ([]models.Answer, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	answers := make([]models.Answer, 0)
	if err = cursor.All(ctx, &answers); err != nil {
		return nil, err
	}

	return answers, nil
}
