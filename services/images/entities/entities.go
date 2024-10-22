package entities

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Taska struct {
	URL     string `validate:"required,url"`
	Caption string `validate:"required,min=1"`
}

type RandImg struct {
	URL     string `validate:"required,url"`
	Caption string `validate:"required,min=1"`
}

func NewTaska(url string) (Taska, error) {
	validate := validator.New()
	taksa := Taska{
		URL:     url,
		Caption: "Собака умная может и самоутилизироваться )\n😍😍😍😍",
	}

	err := validate.Struct(taksa)
	if err != nil {
		return Taska{}, fmt.Errorf("new taksa validation error: %w", err)
	}

	return taksa, nil
}

func NewRandImg(url string) (RandImg, error) {
	validate := validator.New()
	img := RandImg{
		URL:     url,
		Caption: "вообще рандомно:",
	}

	err := validate.Struct(img)
	if err != nil {
		return RandImg{}, fmt.Errorf("new rand img validation error: %w", err)
	}

	return img, nil
}
