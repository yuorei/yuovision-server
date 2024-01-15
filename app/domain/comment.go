package domain

import (
	"fmt"
	"time"

	"github.com/yuorei/video-server/app/driver/db/mongodb/collection"
)

func NewCommentID() string {
	return fmt.Sprintf("%s%s%s", "comment", IDSeparator, NewUUID())
}

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

func NewPostComment(id, videoID, userID, name, text string) *Comment {
	return &Comment{
		ID:        id,
		VideoID:   videoID,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: &User{
			ID:   userID,
			Name: name,
		},
	}
}

func NewCommentForDB(id, videoID, userID, name, text string) *collection.Comment {
	return &collection.Comment{
		ID:        id,
		VideoID:   videoID,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: &collection.User{
			ID:   userID,
			Name: name,
		},
	}
}
