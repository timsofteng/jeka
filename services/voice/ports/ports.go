package ports

import (
	"context"

	"github.com/timsofteng/jeka/services/voice/entities"
)

type RandomText string

type Repo interface {
	Add(ctx context.Context, voiceID string) error
	Rand(ctx context.Context) (voice entities.RandVoice, err error)
	Count(ctx context.Context) (count uint, err error)
}
