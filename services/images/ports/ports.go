package ports

import (
	"context"
	"telegraminput/services/images/entities"
)

type Repo interface {
	Taksa(ctx context.Context) (entities.Taska, error)
	RandImg(ctx context.Context) (entities.RandImg, error)
}
