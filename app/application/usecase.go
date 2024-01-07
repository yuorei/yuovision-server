package application

import (
	"github.com/yuorei/video-server/app/application/port"
)

type UseCase struct {
	port.VideoInputPort
}

func NewUseCase(application *Application) *UseCase {
	return &UseCase{
		VideoInputPort: application,
	}
}
