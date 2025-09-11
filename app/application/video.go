package application

import (
	"context"
	"sort"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type VideoUseCase struct {
	videoRepo port.VideoPort
}

func NewVideoUseCase(videoRepo port.VideoPort) *VideoUseCase {
	return &VideoUseCase{
		videoRepo: videoRepo,
	}
}

func (uc *VideoUseCase) GetVideos(ctx context.Context) ([]*domain.Video, error) {
	videos, err := uc.videoRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[j].CreatedAt.Before(videos[i].CreatedAt)
	})

	return videos, nil
}

func (uc *VideoUseCase) GetVideo(ctx context.Context, videoID string) (*domain.Video, error) {
	return uc.videoRepo.GetByID(ctx, videoID)
}

func (uc *VideoUseCase) GetWatchCount(ctx context.Context, videoID string) (int, error) {
	video, err := uc.videoRepo.GetByID(ctx, videoID)
	if err != nil {
		return 0, err
	}
	return video.WatchCount, nil
}

func (uc *VideoUseCase) IncrementWatchCount(ctx context.Context, videoID, userID string) (int, error) {
	video, err := uc.videoRepo.GetByID(ctx, videoID)
	if err != nil {
		return 0, err
	}

	video.WatchCount++
	err = uc.videoRepo.Update(ctx, video)
	if err != nil {
		return 0, err
	}

	return video.WatchCount, nil
}

func (uc *VideoUseCase) CutVideo(ctx context.Context, videoID string, start, end int) (string, error) {
	// TODO: Implement video cutting logic
	return "", nil
}
