package application

import "github.com/yuorei/video-server/app/application/port"

type VideoUseCase struct {
	videoRepository port.VideoRepository
}

func NewVideoUseCase(videoRepository port.VideoRepository) *VideoUseCase {
	return &VideoUseCase{
		videoRepository: videoRepository,
	}
}
