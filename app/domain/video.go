package domain

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
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
		ID             string
		Video          graphql.Upload
		ThumbnailImage *graphql.Upload
		Title          string
		Description    *string
	}

	UploadVideoResponse struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		CreatedAt         time.Time
	}

	UploadVideoForStorageResponse struct {
		VideoURL          string
		ThumbnailImageURL string
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

func NewUploadVideo(id string, video graphql.Upload, thumbnailImage *graphql.Upload, title string, description *string) *UploadVideo {
	return &UploadVideo{
		ID:             id,
		Video:          video,
		ThumbnailImage: thumbnailImage,
		Title:          title,
		Description:    description,
	}
}

func NewVideoFile(id string, video io.ReadSeeker) *VideoFile {
	return &VideoFile{
		ID:    id,
		Video: video,
	}
}
