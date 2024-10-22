package ports

import "context"

type Services interface {
	RandVideo() (respWithURL string, err error)
	RandImg(ctx context.Context) (url string, err error)
	Taksa(ctx context.Context) (respWithURL string, err error)
}
