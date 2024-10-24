package voice

import (
	"github.com/timsofteng/jeka/services/voice/ports"
)

type Voice struct{ ports.Repo }

func New(repo ports.Repo) *Voice {
	return &Voice{repo}
}
