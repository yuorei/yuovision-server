package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) InsertComment(ctx context.Context, postComment *domain.Comment) (*domain.Comment, error) {
	collection := i.db.Database.Collection("comment")
	if collection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	commentForDB := domain.NewCommentForDB(postComment.ID, postComment.VideoID, postComment.User.ID, postComment.User.Name, postComment.Text)
	insertResult, err := collection.InsertOne(ctx, commentForDB)
	if err != nil {
		return nil, err
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return postComment, nil
}
