package application

import (
	"context"
	"sort"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
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
	videos, err := a.Video.videoRepository.GetVideosFromDB(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[j].CreatedAt.Before(videos[i].CreatedAt)
	})

	return videos, nil
}

func (a *Application) GetVideosByUserID(ctx context.Context, userID string) ([]*domain.Video, error) {
	videos, err := a.Video.videoRepository.GetVideosByUserIDFromDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[j].CreatedAt.Before(videos[i].CreatedAt)
	})

	return videos, nil
}

func (a *Application) GetVideo(ctx context.Context, videoID string) (*domain.Video, error) {
	return a.Video.videoRepository.GetVideoFromDB(ctx, videoID)
}

func (a *Application) UploadVideo(ctx context.Context, video *domain.UploadVideo, userID string, imageURL string) (*domain.UploadVideoResponse, error) {
	videofile := domain.NewVideoFile(video.ID, video.Video)
	err := a.Video.videoRepository.ConvertVideoHLS(ctx, videofile)
	if err != nil {
		return nil, err
	}

	videoURL, err := a.Video.videoRepository.UploadVideoForStorage(ctx, videofile)
	if err != nil {
		return nil, err
	}

	videoResponse, err := a.Video.videoRepository.InsertVideo(ctx, video.ID, videoURL, imageURL, video.Title, video.Description, userID, video.Tags, video.IsAdult, video.IsPrivate, video.IsExternalCutout, video.IsAd)
	if err != nil {
		return nil, err
	}

	return videoResponse, nil
}

func (a *Application) GetWatchCount(ctx context.Context, videoID string) (int, error) {
	return a.Video.videoRepository.GetWatchCount(ctx, videoID)
}

func (a *Application) IncrementWatchCount(ctx context.Context, videoID, userID string) (int, error) {
	ok, err := a.Video.videoRepository.ChechWatchCount(ctx, videoID, userID)
	if err != nil {
		return 0, err
	} else if !ok {
		return 0, nil
	}

	return a.Video.videoRepository.IncrementWatchCount(ctx, videoID, userID)
}

func (a *Application) CutVideo(ctx context.Context, videoID, userID string, start, end int) (string, error) {
	return a.Video.videoRepository.CutVideo(ctx, videoID, userID, start, end)
}
