package video

import "telegraminput/services/video/ports"

type Video struct{ ports.Repo }

func New(repo ports.Repo) *Video {
	return &Video{repo}
}
