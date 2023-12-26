package infrastructure

import "github.com/yuorei/video-server/app/application/port"

type Infrastructure struct {
	port.VideoRepository
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}
