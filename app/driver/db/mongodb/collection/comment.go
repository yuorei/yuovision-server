package collection

import "time"

type (
	Comment struct {
		ID        string `bson:"_id"`
		VideoID   string
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
		User      *User
	}
)
