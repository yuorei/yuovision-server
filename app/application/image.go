package application

import (
	"github.com/yuorei/video-server/app/application/port"
)

type ImageUseCase struct {
	imageRepository port.ImageRepository
}

func NewImageUseCase(imageRepository port.ImageRepository) *ImageUseCase {
	return &ImageUseCase{
		imageRepository: imageRepository,
	}
}
