package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"gogo-form/database"
	"gogo-form/models"
)

type FormRepository struct {
	collection *mongo.Collection
}

func NewFormRepository() *FormRepository {
	collection := database.GetCollection("forms")
	return &FormRepository{collection}
}

func (repo *FormRepository) Create(form models.Form) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if form.Questions == nil {
		form.Questions = []models.Question{}
	}

	return repo.collection.InsertOne(ctx, form)
}

func (repo *FormRepository) GetAll() ([]models.Form, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	forms := make([]models.Form, 0)
	if err = cursor.All(ctx, &forms); err != nil {
		return nil, err
	}

	return forms, nil
}