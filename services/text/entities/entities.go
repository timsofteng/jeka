package entities

import (
	"fmt"

	"github.com/go-playground/validator"
)

type RandText struct {
	Text string `validate:"required,min=1"`
}

func NewRandText(t string) (RandText, error) {
	validate := validator.New()

	text := RandText{Text: t}

	err := validate.Struct(text)
	if err != nil {
		return RandText{}, fmt.Errorf("new text validation error: %w", err)
	}

	return text, nil
}
