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
	videoUseCase := NewVideoUseCase(infra.Video)
	imageUseCase := NewImageUseCase(infra.Image)
	userUseCase := NewUserUseCase(infra.User)
	commentUseCase := NewCommentUseCase(infra.Comment)

	return &Application{
		Video:   videoUseCase,
		Image:   imageUseCase,
		User:    userUseCase,
		Comment: commentUseCase,
	}
}
