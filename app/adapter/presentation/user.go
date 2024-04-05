package presentation

import (
	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func NewUserService(app *application.Application) *UserService {
	return &UserService{
		usecase: application.NewUseCase(app),
	}
}

type UserService struct {
	video_grpc.UnimplementedUserServiceServer
	usecase *application.UseCase
}
