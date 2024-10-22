package entities

import (
	"fmt"

	"github.com/go-playground/validator"
)

type RandVideo struct {
	URL     string `validate:"required,url"`
	Caption string `validate:"required,min=1"`
}

func NewRandVideo(url string) (RandVideo, error) {
	validate := validator.New()
	taksa := RandVideo{
		URL:     url,
		Caption: "взгляните на это видео:",
	}

	err := validate.Struct(taksa)
	if err != nil {
		return RandVideo{}, fmt.Errorf("new taksa validation error: %w", err)
	}

	return taksa, nil
}
