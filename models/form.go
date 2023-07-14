package models

import (
	"gogo-form/domain"

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

func (f *Form) ToEntity() domain.Form {
	questions := make([]domain.Question, len(f.Questions))
	for i, question := range f.Questions {
		questions[i] = domain.Question{
			Text: 		question.Text,
			Type: 		question.Type,
			Options: 	question.Options,
		}
	}

    return domain.Form{
		ID:          f.ID.Hex(),
		Name:        f.Name,
		Description: f.Description,
		Questions:   questions,
	}
}

func (f *Form) FromEntity(e domain.Form) error {
    id, err := primitive.ObjectIDFromHex(e.ID)
    if err != nil {
        return err
    }
	
    f.ID = id
    f.Name = e.Name
    f.Description = e.Description

    f.Questions = make([]Question, len(e.Questions))
    for i, question := range e.Questions {
        f.Questions[i] = Question{
            Text:    question.Text,
            Type:    question.Type,
            Options: question.Options,
        }
    }
    return nil
}
