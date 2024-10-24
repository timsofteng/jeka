package video

import "github.com/timsofteng/jeka/services/video/ports"

type Video struct{ ports.Repo }

func New(repo ports.Repo) *Video {
	return &Video{repo}
}
