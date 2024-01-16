package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/app/driver/db/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *Infrastructure) GetVideoFromDB(ctx context.Context, id string) (*domain.Video, error) {
	mongoCollection := i.db.Database.Collection("video")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	var videoForDB collection.Video
	err := mongoCollection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&videoForDB)
	if err != nil {
		return nil, err
	}

	video := domain.NewVideo(videoForDB.ID, videoForDB.VideoURL, videoForDB.ThumbnailImageURL, videoForDB.Title, videoForDB.Description, videoForDB.UploaderID, videoForDB.CreatedAt)
	return video, nil
}

func (i *Infrastructure) InsertVideo(ctx context.Context, id string, videoURL string, thumbnailImageURL string, title string, description *string, uploaderID string) (*domain.UploadVideoResponse, error) {
	mongoCollection := i.db.Database.Collection("video")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	videoForDB := domain.NewVideoForDB(id, videoURL, thumbnailImageURL, title, description, uploaderID)
	insertResult, err := mongoCollection.InsertOne(ctx, videoForDB)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return &domain.UploadVideoResponse{
		ID:                videoForDB.ID,
		VideoURL:          videoForDB.VideoURL,
		ThumbnailImageURL: videoForDB.ThumbnailImageURL,
		Title:             videoForDB.Title,
		Description:       videoForDB.Description,
		CreatedAt:         videoForDB.CreatedAt,
	}, nil
}
