package presentation

import (
	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func NewCommentService(app *application.Application) *CommentService {
	return &CommentService{
		usecase: application.NewUseCase(app),
	}
}

type CommentService struct {
	video_grpc.UnimplementedCommentServiceServer
	usecase *application.UseCase
}
