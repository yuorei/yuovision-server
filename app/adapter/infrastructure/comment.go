package infrastructure

import (
	"context"
	"database/sql"
	"time"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/db/sqlc"
)

func (i *Infrastructure) GetCommentsByVideoIDFromDB(ctx context.Context, videoID string) ([]*domain.Comment, error) {
	comment, err := i.db.Database.GetVideoComments(ctx, videoID)
	if err != nil {
		return nil, err
	}

	var comments []*domain.Comment
	for _, c := range comment {
		comments = append(comments, domain.NewComment(c.ID, c.VideoID, c.Text, time.Now(), time.Now(), domain.NewUser(c.UserID.String, c.Name, "", []string{}, false, "")))
	}
	return comments, nil
}

func (i *Infrastructure) InsertComment(ctx context.Context, postComment *domain.Comment) (*domain.Comment, error) {
	_, err := i.db.Database.CreateComment(ctx, sqlc.CreateCommentParams{
		ID:      postComment.ID,
		VideoID: postComment.VideoID,
		Text:    postComment.Text,
		UserID: sql.NullString{
			String: postComment.User.ID,
			Valid:  true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return postComment, nil
}
