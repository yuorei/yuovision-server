package application

import (
	"github.com/yuorei/video-server/app/adapter/infrastructure"
)

type Application struct {
	Video   *VideoUseCase
	Image   *ImageUseCase
	User    *UserUseCase
	Comment *CommentUseCase
}

func NewApplication(infra *infrastructure.Infrastructure) *Application {
	videoUseCase := NewVideoUseCase(infra)
	imageUseCase := NewImageUseCase(infra)
	userUseCase := NewUserUseCase(infra)
	CommentUseCase := NewCommentUseCase(infra)

	return &Application{
		Video:   videoUseCase,
		Image:   imageUseCase,
		User:    userUseCase,
		Comment: CommentUseCase,
	}
}
