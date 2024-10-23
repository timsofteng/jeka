package voice

import (
	"telegraminput/services/voice/ports"
)

type Voice struct{ ports.Repo }

func New(repo ports.Repo) *Voice {
	return &Voice{repo}
}
