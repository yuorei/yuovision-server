package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type ImageUseCase struct {
	imageRepository port.ImageRepository
}

func NewImageUseCase(imageRepository port.ImageRepository) *ImageUseCase {
	return &ImageUseCase{
		imageRepository: imageRepository,
	}
}

func (a *Application) UploadThumbnail(ctx context.Context, thumbnail domain.ThumbnailImage) error {
	_, err := a.Image.imageRepository.UploadImageForStorage(ctx, thumbnail.ID)
	if err != nil {
		return err
	}
	return nil
}
