package services

import (
	"context"
	"fmt"
	apperrors "telegraminput/lib/errors"
	"telegraminput/lib/logger"
	"telegraminput/services/httpserver"
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
	logger   logger.Logger
}

func New(logger logger.Logger, services Services) *Adapters {
	return &Adapters{
		services: services,
		logger:   logger,
	}
}

//nolint:ireturn
func (a *Adapters) RandText(
	ctx context.Context,
	_ httpserver.RandTextRequestObject,
) (httpserver.RandTextResponseObject, error) {
	res, err := a.services.Text.Rand(ctx)
	if err != nil {
		return httpserver.RandText500JSONResponse{
			Message: apperrors.ErrInternal.Error(),
		}, fmt.Errorf("failed to call text rand service: %w", err)
	}

	return httpserver.RandText200JSONResponse{Text: res.Text}, nil
}
