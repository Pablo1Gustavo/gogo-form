package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	Text    string   `bson:"text" json:"text,omitempty"`
	Type    string   `bson:"type" json:"type"`
	Options []string `bson:"options,omitempty" json:"options,omitempty"`
}

type Form struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Questions   []Question         `bson:"questions" json:"questions"`
}
