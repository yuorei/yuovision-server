package collection

import "time"

type (
	Video struct {
		ID                string `bson:"_id"`
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		CreatedAt         time.Time
		UpdatedAt         time.Time
		// Tags              []string
		UploaderID string
	}
)

func NewVideoCollection(id string, videoURL string, thumbnailImageURL string, title string, description *string, uploaderID string) *Video {
	return &Video{
		ID:                id,
		VideoURL:          videoURL,
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Description:       description,
		UploaderID:        uploaderID,
		CreatedAt:         time.Now(),
	}
}
