package application

import (
	"github.com/yuorei/video-server/app/adapter/infrastructure"
)

type Application struct {
	Video *VideoUseCase
	User  *UserUseCase
}

func NewApplication() *Application {
	videoUseCase := NewVideoUseCase(infrastructure.NewInfrastructure())
	userUseCase := NewUserUseCase(infrastructure.NewInfrastructure())
	return &Application{
		Video: videoUseCase,
		User:  userUseCase,
	}
}
