package domain

import (
	"fmt"
	"io"
	"time"
)

type (
	VideoProcessingStatus string

	Video struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		UploaderID        string
		Tags              []string
		IsAdult           bool
		IsPrivate         bool
		IsExternalCutout  bool
		IsAd              bool
		CreatedAt         time.Time
		UpdatedAt         time.Time
		WatchCount        int
	}

	VideoProcessingInfo struct {
		ID        string
		VideoID   string
		Status    VideoProcessingStatus
		Progress  int
		Message   *string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UploadVideo struct {
		ID               string
		Video            io.ReadSeeker
		Title            string
		Description      *string
		Tags             []string
		IsAdult          bool
		IsPrivate        bool
		IsExternalCutout bool
		IsAd             bool
	}

	UploadVideoResponse struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		UploaderID        string
		Tags              []string
		IsAdult           bool
		IsPrivate         bool
		IsExternalCutout  bool
		IsAd              bool
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

const (
	VideoProcessingStatusUploaded   VideoProcessingStatus = "UPLOADED"
	VideoProcessingStatusProcessing VideoProcessingStatus = "PROCESSING"
	VideoProcessingStatusCompleted  VideoProcessingStatus = "COMPLETED"
	VideoProcessingStatusFailed     VideoProcessingStatus = "FAILED"
)

func NewVideoID() string {
	return fmt.Sprintf("%s%s%s", "video", IDSeparator, NewUUID())
}

func NewVideoProcessingInfoID() string {
	return fmt.Sprintf("%s%s%s", "video_processing", IDSeparator, NewUUID())
}

func NewVideo(id string, videoURL string, thumbnailImageURL string, title string, description *string, tags []string, watchCount int, private bool, adult bool, externalCutout bool, isAd bool, uploaderID string, createdAt time.Time, updatedAt time.Time) *Video {
	return &Video{
		ID:                id,
		VideoURL:          videoURL,
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Description:       description,
		Tags:              tags,
		IsPrivate:         private,
		IsAdult:           adult,
		IsExternalCutout:  externalCutout,
		IsAd:              isAd,
		UploaderID:        uploaderID,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
		WatchCount:        watchCount,
	}
}

func NewUploadVideo(id string, video io.ReadSeeker, title string, description *string, tags []string, isAdult, isPrivate, isExternalCutout, isAd bool) *UploadVideo {
	return &UploadVideo{
		ID:               id,
		Video:            video,
		Title:            title,
		Description:      description,
		Tags:             tags,
		IsAdult:          isAdult,
		IsPrivate:        isPrivate,
		IsExternalCutout: isExternalCutout,
		IsAd:             isAd,
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

func NewVideoProcessingInfo(id, videoID string, status VideoProcessingStatus, progress int, message *string, createdAt, updatedAt time.Time) *VideoProcessingInfo {
	return &VideoProcessingInfo{
		ID:        id,
		VideoID:   videoID,
		Status:    status,
		Progress:  progress,
		Message:   message,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
