package presentation

import (
	"context"

	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *VideoService) Video(ctx context.Context, id *video_grpc.VideoID) (*video_grpc.VideoPayload, error) {
	video, err := s.usecase.GetVideo(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	return &video_grpc.VideoPayload{
		Id:                video.ID,
		VideoUrl:          video.VideoURL,
		Title:             video.Title,
		ThumbnailImageUrl: video.ThumbnailImageURL,
		Description:       *video.Description,
		CreatedAt:         timestamppb.New(video.CreatedAt),
		UpdatedAt:         timestamppb.New(video.UpdatedAt),
		UserId:            video.UploaderID,
	}, nil
}
