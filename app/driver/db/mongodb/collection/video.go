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
