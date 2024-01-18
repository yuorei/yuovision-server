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
		ID               string
		Video            io.ReadSeeker
		VideoContentType string
		ThumbnailImage   *io.ReadSeeker
		ImageContentType string
		Title            string
		Description      *string
	}

	UploadVideoResponse struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		CreatedAt         time.Time
	}

	VideoFile struct {
		ID    string
		Video io.ReadSeeker
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

func NewUploadVideo(id string, video io.ReadSeeker, videoContentType string, thumbnailImage *io.ReadSeeker, imageContentType string, title string, description *string) *UploadVideo {
	return &UploadVideo{
		ID:               id,
		Video:            video,
		VideoContentType: videoContentType,
		ThumbnailImage:   thumbnailImage,
		ImageContentType: imageContentType,
		Title:            title,
		Description:      description,
	}
}

func NewVideoFile(id string, video io.ReadSeeker) *VideoFile {
	return &VideoFile{
		ID:    id,
		Video: video,
	}
}
