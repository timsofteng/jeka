package services

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/timsofteng/jeka/services/images"
	"github.com/timsofteng/jeka/services/text"
	"github.com/timsofteng/jeka/services/video"
	"github.com/timsofteng/jeka/services/voice"

	tele "gopkg.in/telebot.v4"
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

func (a *Adapters) Rand(ctx context.Context) (any, error) {
	//nolint:gosec,mnd
	if rand.IntN(100) < 85 {
		return a.RandText(ctx)
	}

	return a.RandVoice(ctx)
}

func (a *Adapters) RandVideo(ctx context.Context) (string, error) {
	resp, err := a.services.Video.RandVideo(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call Video.RandVideo: %w", err)
	}

	return fmt.Sprintf("%s:\n\n%s", resp.Caption, resp.URL), nil
}

func (a *Adapters) RandImg(ctx context.Context) (*tele.Photo, error) {
	img, err := a.services.Image.RandImg(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call Image.RandImg: %w", err)
	}

	photo := &tele.Photo{
		File: tele.File{FileURL: img.URL}, Caption: img.Caption,
	}

	return photo, nil
}

func (a *Adapters) Taksa(ctx context.Context) (*tele.Photo, error) {
	img, err := a.services.Image.Taksa(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call Image.Taska: %w", err)
	}

	photo := &tele.Photo{
		File: tele.File{FileURL: img.URL}, Caption: img.Caption,
	}

	return photo, nil
}

func (a *Adapters) RandText(ctx context.Context) (string, error) {
	res, err := a.services.Text.Rand(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to call randtext: %w", err)
	}

	return res.Text, nil
}

func (a *Adapters) RandVoice(ctx context.Context) (*tele.Voice, error) {
	res, err := a.services.Voice.Rand(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call rand voice: %w", err)
	}

	return &tele.Voice{File: tele.File{FileID: res.ID}}, nil
}
