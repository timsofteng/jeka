package ports

import (
	"context"

	tele "gopkg.in/telebot.v4"
)

type Services interface {
	Rand(ctx context.Context) (any, error)
	RandVideo(ctx context.Context) (*tele.Video, error)
	RandImg(ctx context.Context) (*tele.Photo, error)
	Taksa(ctx context.Context) (*tele.Photo, error)
	RandText(ctx context.Context) (text string, err error)
	RandVoice(ctx context.Context) (voice *tele.Voice, err error)
}
