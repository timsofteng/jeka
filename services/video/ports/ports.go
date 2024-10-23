package ports

import (
	"context"
	"telegraminput/services/video/entities"
)

type Repo interface {
	RandVideo(ctx context.Context) (entities.RandVideo, error)
}
