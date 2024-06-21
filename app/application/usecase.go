package application

import (
	"github.com/yuorei/video-server/app/application/port"
)

type UseCase struct {
	port.VideoInputPort
	port.UserInputPort
	port.CommentInputPort
	port.ImageInputPort
}

func NewUseCase(application *Application) *UseCase {
	return &UseCase{
		VideoInputPort:   application,
		UserInputPort:    application,
		CommentInputPort: application,
		ImageInputPort:   application,
	}
}
