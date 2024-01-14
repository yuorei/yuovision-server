package application

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/middleware"
)

func (a *Application) UploadVideo(ctx context.Context, video *domain.UploadVideo) (*domain.UploadVideoResponse, error) {
	id, err := middleware.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	videofile := domain.NewVideoFile(video.ID, video.Video.File)

	err = a.Video.videoRepository.ConvertVideoHLS(ctx, videofile)
	if err != nil {
		return nil, err
	}

	uploadVideoForStorageResponse, err := a.Video.videoRepository.UploadVideoForStorage(ctx, videofile)
	if err != nil {
		return nil, err
	}

	videoResponse, err := a.Video.videoRepository.InsertVideo(ctx, video.ID, uploadVideoForStorageResponse.VideoURL, uploadVideoForStorageResponse.ThumbnailImageURL, video.Title, video.Description, id)
	if err != nil {
		return nil, err
	}

	return videoResponse, nil
}
