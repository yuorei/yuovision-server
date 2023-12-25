package application

import (
	"time"

	"github.com/yuorei/video-server/app/domain"
)

func (a *Application) UploadVideo(video *domain.UploadVideo) (*domain.UploadVideoResponse, error) {
	videofile := domain.NewVideoFile(video.ID, video.Video.File)

	err := a.Video.videoRepository.ConvertVideoHLS(videofile)
	if err != nil {
		return nil, err
	}

	uploadVideoForStorageResponse, err := a.Video.videoRepository.UploadVideoForStorage(videofile)
	if err != nil {
		return nil, err
	}

	return &domain.UploadVideoResponse{
		ID:                 video.ID,
		VideoURL:           uploadVideoForStorageResponse.VideoURL,
		VideoSize:          1000,
		ThumbnailImageURL:  uploadVideoForStorageResponse.ThumbnailImageURL,
		ThumbnailImageSize: 1000,
		Title:              video.Title,
		Description:        video.Description,
		CreatedAt:          time.Now(),
	}, nil
}
