package models

import (
	"gogo-form/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FormID     primitive.ObjectID `bson:"form_id" json:"form_id"`
	AnsweredAt time.Time          `bson:"answered_at" json:"answered_at"`
	Answers    []interface{}      `bson:"answers" json:"answers"`
}

func (a *Answer) ToEntity() domain.Answer {
	return domain.Answer{
		ID:         a.ID.Hex(),
		FormID:     a.FormID.Hex(),
		AnsweredAt: a.AnsweredAt,
		Answers:    a.Answers,
	}
}

func (a *Answer) FromEntity(entity domain.Answer) error {
	id, err := primitive.ObjectIDFromHex(entity.ID)
	if err != nil {
		return err
	}
	formID, err := primitive.ObjectIDFromHex(entity.FormID)
	if err != nil {
		return err
	}

	a.ID = id
	a.FormID = formID
	a.AnsweredAt = entity.AnsweredAt
	a.Answers = entity.Answers

	return nil
}
