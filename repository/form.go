package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"gogo-form/database"
	"gogo-form/domain"
	"gogo-form/models"
)

type FormRepository struct {
	collection *mongo.Collection
}

func NewFormRepository() domain.FormRepository {
	collection := database.GetCollection("forms")
	return &FormRepository{collection}
}

func (r *FormRepository) Create(ctx context.Context, form domain.Form) (domain.Form, error) {
	form.ID = primitive.NewObjectID().Hex()

	modelForm := models.Form{}

	if err := modelForm.FromEntity(form); err != nil {
		return form, err
	}

	_, err := r.collection.InsertOne(ctx, modelForm)
	if err != nil {
		return form, err
	}

	return form, nil
}

func (r *FormRepository) GetAll(ctx context.Context, name string) ([]domain.Form, error) {
	filter := bson.M{}
	if name != "" {
		filter["name"] = primitive.Regex{Pattern: name, Options: "i"}
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	forms := make([]domain.Form, 0)
	var modelForm models.Form

	for cursor.Next(ctx) {
		err := cursor.Decode(&modelForm)
		if err != nil {
			return nil, err
		}
		forms = append(forms, modelForm.ToEntity())
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return forms, nil
}

func (r *FormRepository) GetOne(ctx context.Context, id string) (domain.Form, error) {
	var form models.Form

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Form{}, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&form)
	if err != nil {
		return domain.Form{}, err
	}

	return form.ToEntity(), nil
}

func (r *FormRepository) Update(ctx context.Context, form domain.Form, id string) (domain.Form, error) {
	form.ID = id
	modelForm := models.Form{}

	if err := modelForm.FromEntity(form); err != nil {
		return form, err
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Form{}, err
	}

	updateResult := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.M{"$set": modelForm})

	if updateResult.Err() != nil {
		return domain.Form{}, updateResult.Err()
	}

	return form, nil
}

func (r *FormRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, _ := r.collection.DeleteOne(ctx, bson.M{"_id": objID})

	if result.DeletedCount == 0 {
		return errors.New("not found")
	}

	return nil
}
