package images

import (
	"github.com/timsofteng/jeka/services/images/ports"
)

type Images struct {
	ports.Repo
}

func New(repo ports.Repo) *Images {
	return &Images{repo}
}
