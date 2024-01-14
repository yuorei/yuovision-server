package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) InsertVideo(ctx context.Context, id string, videoURL string, thumbnailImageURL string, title string, description *string, uploaderID string) (*domain.UploadVideoResponse, error) {
	collection := i.db.Database.Collection("video")
	if collection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	videoForDB := domain.NewVideoForDB(id, videoURL, thumbnailImageURL, title, description, uploaderID)
	insertResult, err := collection.InsertOne(ctx, videoForDB)
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
