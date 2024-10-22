package ports

import (
	"telegraminput/services/video/entities"
)

type Repo interface {
	RandVideo() (entities.RandVideo, error)
}
