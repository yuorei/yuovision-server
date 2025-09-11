package application

import (
	"context"
	"io"

	"github.com/yuorei/video-server/app/application/port"
)

type ImageUseCase struct {
	imageRepo port.ImagePort
}

func NewImageUseCase(imageRepo port.ImagePort) *ImageUseCase {
	return &ImageUseCase{
		imageRepo: imageRepo,
	}
}

func (uc *ImageUseCase) UploadImage(ctx context.Context, key string, file io.Reader, contentType string) (string, error) {
	return uc.imageRepo.Upload(ctx, key, file, contentType)
}

func (uc *ImageUseCase) GetPresignedURL(ctx context.Context, key string) (string, error) {
	return uc.imageRepo.GetPresignedURL(ctx, key)
}

func (uc *ImageUseCase) GetUploadURL(ctx context.Context, key string, contentType string) (string, error) {
	return uc.imageRepo.GetUploadURL(ctx, key, contentType)
}
