package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	ID          primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	FormID  	primitive.ObjectID 	`bson:"form_id" json:"form_id"`
	AnsweredAt	time.Time          	`bson:"answered_at" json:"answered_at"`
	Answers     []interface{}		`bson:"answers" json:"answers"`
}
