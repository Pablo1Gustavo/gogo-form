package domain

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
