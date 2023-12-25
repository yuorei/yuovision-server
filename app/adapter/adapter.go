package adapter

import (
	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/app/application/port"
)

type Adapter struct {
	port.VideoInputPort
}

func NewAdapter(application *application.Application) *Adapter {
	return &Adapter{
		VideoInputPort: application,
	}
}
