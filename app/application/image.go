package application

import (
	"context"
	"os"

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
	videoID := thumbnail.ID
	if thumbnail.ContentType != "" {
		thumbnailImage, err := os.Open(videoID + "." + thumbnail.ContentType)
		if err != nil {
			return err
		}

		_, err = a.Image.imageRepository.ConvertThumbnailToWebp(ctx, thumbnailImage, thumbnail.ContentType, videoID)
		if err != nil {
			return err
		}
	} else {
		err := a.Image.imageRepository.CreateThumbnail(ctx, videoID)
		if err != nil {
			return err
		}
	}

	thumbnail = domain.NewThumbnailImage(videoID, thumbnail.ContentType)
	_, err := a.Image.imageRepository.UploadImageForStorage(ctx, videoID)
	if err != nil {
		return err
	}

	return nil
}
