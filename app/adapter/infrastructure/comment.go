package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/app/driver/db/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *Infrastructure) GetCommentsByVideoIDFromDB(ctx context.Context, videoID string) ([]*domain.Comment, error) {
	mongoCollection := i.db.Database.Collection("comment")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	cursor, err := mongoCollection.Find(ctx, bson.D{{"videoid", videoID}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx) {
		var commentForDB collection.Comment
		if err := cursor.Decode(&commentForDB); err != nil {
			return nil, err
		}
		comment := domain.NewComment(commentForDB.ID, videoID, commentForDB.Text, commentForDB.CreatedAt, commentForDB.UpdatedAt, domain.NewUser(commentForDB.User.ID, commentForDB.User.Name, commentForDB.User.ProfileImageURL, commentForDB.User.SubscribeChannelIDs))
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

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
