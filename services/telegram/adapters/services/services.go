package services

import (
	"context"
	"fmt"
	"telegraminput/services/images"
	"telegraminput/services/text"
	"telegraminput/services/video"
	"telegraminput/services/voice"
)

type Services struct {
	Video *video.Video
	Image *images.Images
	Text  *text.Text
	Voice *voice.Voice
}

type Adapters struct {
	services Services
}

func New(services Services) *Adapters {
	return &Adapters{
		services: services,
	}
}

func (a *Adapters) RandVideo(ctx context.Context) (string, error) {
	video, err := a.services.Video.RandVideo(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call Video.RandVideo: %w", err)
	}

	return fmt.Sprintf("%s \n %s", video.Caption, video.URL), nil
}

func (a *Adapters) RandImg(ctx context.Context) (string, string, error) {
	img, err := a.services.Image.RandImg(ctx)
	if err != nil {
		return "", "", fmt.Errorf("failed to call Image.RandImg: %w", err)
	}

	return img.URL, img.Caption, nil
}

func (a *Adapters) Taksa(ctx context.Context) (string, string, error) {
	taksa, err := a.services.Image.Taksa(ctx)
	if err != nil {
		return "", "", fmt.Errorf("failed to call Image.Taska: %w", err)
	}

	return taksa.URL, taksa.Caption, nil
}

func (a *Adapters) RandText(ctx context.Context) (string, error) {
	res, err := a.services.Text.Rand(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call randtext: %w", err)
	}

	return res.Text, nil
}

func (a *Adapters) RandVoice(ctx context.Context) (string, error) {
	res, err := a.services.Voice.Rand(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call rand voice: %w", err)
	}

	return res.ID, nil
}
