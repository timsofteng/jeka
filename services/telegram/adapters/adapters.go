package adapters

import (
	"context"
	"fmt"
	"telegraminput/services/images"
	"telegraminput/services/video"
)

type Services struct {
	Video *video.Video
	Image *images.Images
}

type Adapters struct {
	services Services
}

func New(services Services) *Adapters {
	return &Adapters{
		services: services,
	}
}

func (a *Adapters) RandVideo() (string, error) {
	video, err := a.services.Video.RandVideo()
	if err != nil {
		return "", fmt.Errorf("failed to call Video.RandVideo: %w", err)
	}

	return fmt.Sprintf("%s \n %s", video.Caption, video.URL), nil
}

func (a *Adapters) RandImg(ctx context.Context) (string, error) {
	img, err := a.services.Image.RandImg(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call Image.RandImg: %w", err)
	}

	return fmt.Sprintf("%s \n %s", img.Caption, img.URL), nil
}

func (a *Adapters) Taksa(ctx context.Context) (string, error) {
	taksa, err := a.services.Image.Taksa(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call Image.Taska: %w", err)
	}

	return fmt.Sprintf("%s \n %s", taksa.Caption, taksa.URL), nil
}
