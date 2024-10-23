package text

import (
	"telegraminput/services/text/ports"
)

type Text struct{ ports.Repo }

func New(repo ports.Repo) *Text {
	return &Text{repo}
}
