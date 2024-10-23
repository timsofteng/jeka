package ports

import (
	"context"
	"telegraminput/services/text/entities"
)

type RandomText string

type Repo interface {
	Add(ctx context.Context, text string) error
	Rand(ctx context.Context) (text entities.RandText, err error)
	Count(ctx context.Context) (count uint, err error)
}
