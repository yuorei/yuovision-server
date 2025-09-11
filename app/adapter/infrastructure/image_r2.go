package infrastructure

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/yuorei/video-server/app/driver/storage"
)

type ImageR2Repository struct {
	r2Client *storage.R2Client
}

func NewImageRepository(r2Client *storage.R2Client) *ImageR2Repository {
	return &ImageR2Repository{
		r2Client: r2Client,
	}
}

func (r *ImageR2Repository) Upload(ctx context.Context, key string, file io.Reader, contentType string) (string, error) {
	err := r.r2Client.UploadFile(ctx, key, file, contentType)
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	// Generate presigned URL for access
	url, err := r.r2Client.GetPresignedURL(ctx, key, time.Hour*24*30) // 30 days expiration
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url, nil
}

func (r *ImageR2Repository) GetPresignedURL(ctx context.Context, key string) (string, error) {
	url, err := r.r2Client.GetPresignedURL(ctx, key, time.Hour*24) // 24 hours expiration
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL for %s: %w", key, err)
	}

	return url, nil
}

func (r *ImageR2Repository) GetUploadURL(ctx context.Context, key string, contentType string) (string, error) {
	url, err := r.r2Client.GetUploadPresignedURL(ctx, key, contentType, time.Minute*15) // 15 minutes for upload
	if err != nil {
		return "", fmt.Errorf("failed to generate upload URL for %s: %w", key, err)
	}

	return url, nil
}

func (r *ImageR2Repository) Delete(ctx context.Context, key string) error {
	err := r.r2Client.DeleteFile(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete image %s: %w", key, err)
	}

	return nil
}
