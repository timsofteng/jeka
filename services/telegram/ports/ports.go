package ports

import "context"

type Services interface {
	RandVideo(ctx context.Context) (respWithURL string, err error)
	RandImg(ctx context.Context) (url string, caption string, err error)
	Taksa(ctx context.Context) (url string, caption string, err error)
	RandText(ctx context.Context) (text string, err error)
	RandVoice(ctx context.Context) (id string, err error)
}
