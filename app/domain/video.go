package domain

import (
	"fmt"
	"io"
	"time"
)

type (
	Video struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		UploaderID        string
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}

	UploadVideo struct {
		ID          string
		Video       io.ReadSeeker
		Title       string
		Description *string
	}

	UploadVideoResponse struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		UploaderID        string
		CreatedAt         time.Time
	}

	VideoFile struct {
		ID    string
		Video io.ReadSeeker
	}

	ThumbnailImage struct {
		ID          string
		ContentType string
	}
)

func NewVideoID() string {
	return fmt.Sprintf("%s%s%s", "video", IDSeparator, NewUUID())
}

func NewVideo(id string, videoURL string, thumbnailImageURL string, title string, description *string, uploaderID string, createdAt time.Time) *Video {
	return &Video{
		ID:                id,
		VideoURL:          videoURL,
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Description:       description,
		UploaderID:        uploaderID,
		CreatedAt:         createdAt,
	}
}

func NewUploadVideo(id string, video io.ReadSeeker, title string, description *string) *UploadVideo {
	return &UploadVideo{
		ID:          id,
		Video:       video,
		Title:       title,
		Description: description,
	}
}

func NewVideoFile(id string, video io.ReadSeeker) *VideoFile {
	return &VideoFile{
		ID:    id,
		Video: video,
	}
}

func NewThumbnailImage(id, contentType string) ThumbnailImage {
	return ThumbnailImage{
		ID:          id,
		ContentType: contentType,
	}
}
