package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"gogo-form/database"
	"gogo-form/domain"
	"gogo-form/models"
)

type AnswerRepository struct {
	collection *mongo.Collection
}

func NewAnswerRepository() *AnswerRepository {
	collection := database.GetCollection("answers")
	return &AnswerRepository{collection}
}

func (repo *AnswerRepository) Create(ctx context.Context, answer domain.Answer) (domain.Answer, error) {
	answer.ID = primitive.NewObjectID().Hex()

	modelAnswer := new(models.Answer)
	if err := modelAnswer.FromEntity(answer); err != nil {
		return answer, err
	}

	_, err := repo.collection.InsertOne(ctx, modelAnswer)
	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (repo *AnswerRepository) GetAll(ctx context.Context) ([]domain.Answer, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	answers := make([]domain.Answer, 0)
	var modelAnswer models.Answer

	for cursor.Next(ctx) {
		err := cursor.Decode(&modelAnswer)
		if err != nil {
			return nil, err
		}
		answers = append(answers, modelAnswer.ToEntity())
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return answers, nil
}

func (repo *AnswerRepository) GetOne(ctx context.Context, id string) (domain.Answer, error) {
	var answer models.Answer

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Answer{}, err
	}

	err = repo.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&answer)
	if err != nil {
		return domain.Answer{}, err
	}

	return answer.ToEntity(), nil
}
