package entities

import (
	"fmt"

	"github.com/go-playground/validator"
)

type RandVoice struct {
	ID string `validate:"required,min=1"`
}

func NewRandVoice(voiceID string) (RandVoice, error) {
	validate := validator.New()

	text := RandVoice{ID: voiceID}

	err := validate.Struct(text)
	if err != nil {
		return RandVoice{}, fmt.Errorf("new voice validation error: %w", err)
	}

	return text, nil
}
