package application

import (
	"github.com/yuorei/video-server/app/adapter/infrastructure"
)

type Application struct {
	Video *VideoUseCase
}

func NewApplication() *Application {
	videoUseCase := NewVideoUseCase(infrastructure.NewInfrastructure())
	return &Application{
		Video: videoUseCase,
	}
}
