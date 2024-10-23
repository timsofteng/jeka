package ports

import (
	"context"

	tele "gopkg.in/telebot.v4"
)

type Services interface {
	Rand(ctx context.Context) (any, error)
	RandVideo(ctx context.Context) (respWithURL string, err error)
	RandImg(ctx context.Context) (url string, caption string, err error)
	Taksa(ctx context.Context) (url string, caption string, err error)
	RandText(ctx context.Context) (text string, err error)
	RandVoice(ctx context.Context) (voice *tele.Voice, err error)
}
