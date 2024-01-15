package domain

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/yuorei/video-server/app/driver/db/mongo/collection"
)

type (
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

func NewVideoForDB(id string, videoURL string, thumbnailImageURL string, title string, description *string, uploaderID string) *collection.Video {
	return &collection.Video{
		ID:                id,
		VideoURL:          videoURL,
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Description:       description,
		UploaderID:        uploaderID,
		CreatedAt:         time.Now(),
	}
}
