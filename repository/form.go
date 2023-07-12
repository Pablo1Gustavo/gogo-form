package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *FormRepository) Create(ctx context.Context, form models.Form) (*mongo.InsertOneResult, error) {
	if form.Questions == nil {
		form.Questions = []models.Question{}
	}
	return repo.collection.InsertOne(ctx, form)
}

func (repo *FormRepository) GetAll(ctx context.Context) ([]models.Form, error) {
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

func (repo *FormRepository) GetOne(ctx context.Context, id primitive.ObjectID) (*models.Form, error) {
	var form models.Form

	if err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&form); err != nil {
		return nil, err
	}

	return &form, nil
}

