package images

import (
	"telegraminput/services/images/ports"
)

type Images struct {
	ports.Repo
}

func New(repo ports.Repo) *Images {
	return &Images{repo}
}
