package unsplash

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/timsofteng/jeka/services/images/entities"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type resp struct {
	Urls struct {
		Regular string `json:"regular" validate:"required"`
	} `json:"urls" validate:"required"`

	ID string `json:"id" validate:"required"`
}

type Unsplash struct {
	validate   *validator.Validate
	httpClient httpClient
}

func New(unsplashClientID string) *Unsplash {
	validate := validator.New()

	return &Unsplash{
		httpClient: newHTTPClient(unsplashClientID),
		validate:   validate,
	}
}

func (c Unsplash) Taksa(ctx context.Context) (entities.Taska, error) {
	const query = "dachshund"

	url, err := c.randomImg(ctx, query)
	if err != nil {
		return entities.Taska{}, fmt.Errorf("failed to get random img: %w", err)
	}

	taksa, err := entities.NewTaska(url)
	if err != nil {
		return entities.Taska{}, fmt.Errorf("failed to create new taska: %w", err)
	}

	return taksa, nil
}

func (c Unsplash) RandImg(ctx context.Context) (entities.RandImg, error) {
	url, err := c.randomImg(ctx, "")
	if err != nil {
		return entities.RandImg{}, fmt.Errorf("failed to get random img: %w", err)
	}

	img, err := entities.NewRandImg(url)
	if err != nil {
		return entities.RandImg{}, fmt.Errorf("failed to create new taska: %w", err)
	}

	return img, nil
}

func (c Unsplash) randomImg(ctx context.Context,
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
