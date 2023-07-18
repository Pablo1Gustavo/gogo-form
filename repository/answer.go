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

func NewAnswerRepository() domain.AnswerRepository {
	collection := database.GetCollection("answers")
	return &AnswerRepository{collection}
}

func (r *AnswerRepository) Create(ctx context.Context, answer domain.Answer) (domain.Answer, error) {
	answer.ID = primitive.NewObjectID().Hex()

	modelAnswer := models.Answer{}
	modelAnswer.FromEntity(answer)

	_, err := r.collection.InsertOne(ctx, modelAnswer)
	if err != nil {
		return answer, &domain.RequestError{Err: err}
	}

	return answer, nil
}

func (r *AnswerRepository) GetAll(ctx context.Context, formId string) ([]domain.Answer, error) {
	filter := bson.M{}
	if formId != "" {
		objFormId, _ := primitive.ObjectIDFromHex(formId)
		filter["form_id"] = objFormId
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, &domain.RequestError{Err: err}
	}
	defer cursor.Close(ctx)

	answers := make([]domain.Answer, 0)
	var modelAnswer models.Answer

	for cursor.Next(ctx) {
		err := cursor.Decode(&modelAnswer)
		if err != nil {
			return nil, &domain.RequestError{Err: err}
		}
		answers = append(answers, modelAnswer.ToEntity())
	}

	if err := cursor.Err(); err != nil {
		return nil, &domain.RequestError{Err: err}
	}

	return answers, nil
}

func (r *AnswerRepository) GetOne(ctx context.Context, id string) (domain.Answer, error) {
	var answer models.Answer

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Answer{}, &domain.RequestError{Code: 404, Err: err}
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&answer)
	if err != nil {
		return domain.Answer{}, &domain.RequestError{Code: 404, Err: err}
	}

	return answer.ToEntity(), nil
}

func (r *AnswerRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.RequestError{Code: 404, Err: err}
	}

	result, _ := r.collection.DeleteOne(ctx, bson.M{"_id": objID})

	if result.DeletedCount == 0 {
		return &domain.RequestError{Code: 404, Err: err}
	}

	return nil
}
