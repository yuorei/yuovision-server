package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/middleware"
)

type VideoUseCase struct {
	videoRepository port.VideoRepository
}

func NewVideoUseCase(videoRepository port.VideoRepository) *VideoUseCase {
	return &VideoUseCase{
		videoRepository: videoRepository,
	}
}

func (a *Application) GetVideos(ctx context.Context) ([]*domain.Video, error) {
	return a.Video.videoRepository.GetVideosFromDB(ctx)
}

func (a *Application) GetVideo(ctx context.Context, videoID string) (*domain.Video, error) {
	return a.Video.videoRepository.GetVideoFromDB(ctx, videoID)
}

func (a *Application) UploadVideo(ctx context.Context, video *domain.UploadVideo) (*domain.UploadVideoResponse, error) {
	id, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	imageBuffer, err := a.Image.imageRepository.ConvertThumbnailToWebp(ctx, video.ThumbnailImage, video.ImageContentType, video.ID)
	if err != nil {
		return nil, err
	}

	var imageURL string
	if imageBuffer != nil {
		imageURL, err = a.Image.imageRepository.UploadImageForStorage(ctx, video.ID, imageBuffer)
		if err != nil {
			return nil, err
		}
	}

	videofile := domain.NewVideoFile(video.ID, video.Video)
	err = a.Video.videoRepository.ConvertVideoHLS(ctx, videofile)
	if err != nil {
		return nil, err
	}

	videoURL, err := a.Video.videoRepository.UploadVideoForStorage(ctx, videofile)
	if err != nil {
		return nil, err
	}

	videoResponse, err := a.Video.videoRepository.InsertVideo(ctx, video.ID, videoURL, imageURL, video.Title, video.Description, id)
	if err != nil {
		return nil, err
	}

	return videoResponse, nil
}
