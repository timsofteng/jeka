package images

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type resp struct {
	Urls struct {
		Regular string `json:"regular" validate:"required"`
	} `json:"urls" validate:"required"`

	ID string `json:"id" validate:"required"`
}

type Images struct {
	validate   *validator.Validate
	httpClient httpClient
}

func New(unsplashClientID string) *Images {
	validate := validator.New()

	return &Images{
		httpClient: newHTTPClient(unsplashClientID),
		validate:   validate,
	}
}

func (c Images) Taksa(ctx context.Context) (string, string, error) {
	const (
		caption = "Собака умная может и самоутилизироваться )\n😍😍😍😍"
		query   = "dachshund"
	)

	url, err := c.randomImg(ctx, query)
	if err != nil {
		return "", "", fmt.Errorf("failed to get random img: %w", err)
	}

	return url, caption, nil
}

func (c Images) RandomImg(ctx context.Context) (string, string, error) {
	const (
		caption = "вообще рандомно:"
	)

	url, err := c.randomImg(ctx, "")
	if err != nil {
		return "", "", fmt.Errorf("failed to get random img: %w", err)
	}

	return url, caption, nil
}

func (c Images) randomImg(ctx context.Context,
	query string,
) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, "/photos/random", nil)
	if err != nil {
		return "", fmt.Errorf("failed to construct request: %w", err)
	}

	if len(query) > 0 {
		q := req.URL.Query()
		q.Add("query", query)
		req.URL.RawQuery = q.Encode()
	}

	jsonBin, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	var data resp

	err = json.Unmarshal(jsonBin, &data)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal resp: %w", err)
	}

	err = c.validate.Struct(data)
	if err != nil {
		return "", fmt.Errorf("failed to validate resp: %w", err)
	}

	url := data.Urls.Regular

	return url, nil
}
