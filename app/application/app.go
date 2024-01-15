package application

import (
	"github.com/yuorei/video-server/app/adapter/infrastructure"
)

type Application struct {
	Video   *VideoUseCase
	User    *UserUseCase
	Comment *CommentUseCase
}

func NewApplication() *Application {
	infra := infrastructure.NewInfrastructure()

	videoUseCase := NewVideoUseCase(infra)
	userUseCase := NewUserUseCase(infra)
	CommentUseCase := NewCommentUseCase(infra)

	return &Application{
		Video:   videoUseCase,
		User:    userUseCase,
		Comment: CommentUseCase,
	}
}
