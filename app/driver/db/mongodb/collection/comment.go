package collection

import "time"

type (
	Comment struct {
		ID        string
		VideoID   string
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
		User      *User
	}
)
