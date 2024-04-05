package presentation

import (
	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func NewVideoService(app *application.Application) *VideoService {
	return &VideoService{
		usecase: application.NewUseCase(app),
	}
}

type VideoService struct {
	video_grpc.UnimplementedVideoServiceServer
	usecase *application.UseCase
}
