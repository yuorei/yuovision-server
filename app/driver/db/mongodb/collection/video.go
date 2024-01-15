package collection

import "time"

type (
	Video struct {
		ID                string
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
