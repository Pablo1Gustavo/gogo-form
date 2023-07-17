package domain

import (
	"context"
	"time"
)

type Answer struct {
	ID         string        `json:"id,omitempty"`
	FormID     string        `json:"form_id"`
	AnsweredAt time.Time     `json:"answered_at"`
	Answers    []interface{} `json:"answers"`
}

type AnswerRepository interface {
	Create(ctx context.Context, answer Answer) (Answer, error)
	GetAll(ctx context.Context) ([]Answer, error)
	GetOne(ctx context.Context, id string) (Answer, error)
	Delete(ctx context.Context, id string) error
}

func (a Answer) CompatibleWithForm(form Form) bool {
	if len(a.Answers) != len(form.Questions) {
		return false
	}
	for i := 0; i < len(form.Questions); i++ {
		switch form.Questions[i].Type {
		case "text":
			if _, ok := a.Answers[i].(string); !ok {
				return false
			}

		case "boolean":
			if _, ok := a.Answers[i].(bool); !ok {
				return false
			}

		case "option":
			answer, ok := a.Answers[i].(string)
			if !ok {
				return false
			}
			ok = false
			for _, option := range form.Questions[i].Options {
				if answer == option {
					ok = true
					break
				}
			}
			if !ok {
				return false
			}

		default:
			return false
		}
	}
	return true
}
