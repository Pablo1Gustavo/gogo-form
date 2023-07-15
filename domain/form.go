package domain

import "context"

type Question struct {
	Text    string   `json:"text" validate:"required"`
	Type    string   `json:"type" validate:"required,oneof=text boolean option"`
	Options []string `json:"options,omitempty" validate:"required_if=Type option"`
}

type Form struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Questions   []Question `json:"questions" validate:"required,dive"`
}

type FormRepository interface {
	Create(ctx context.Context, form Form) (Form, error)
	GetAll(ctx context.Context, name string) ([]Form, error)
	GetOne(ctx context.Context, id string) (Form, error)
	Update(ctx context.Context, form Form, id string) (Form, error)
	Delete(ctx context.Context, id string) error
}
