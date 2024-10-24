package ports

import (
	"context"

	"github.com/timsofteng/jeka/services/video/entities"
)

type Repo interface {
	RandVideo(ctx context.Context) (entities.RandVideo, error)
}
