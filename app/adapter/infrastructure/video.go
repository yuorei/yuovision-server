package infrastructure

import (
	"context"
	"database/sql"
	"time"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/db/sqlc"
)

func (i *Infrastructure) GetVideosFromDB(ctx context.Context) ([]*domain.Video, error) {
	var videos []*domain.Video
	dbVideos, err := i.db.Database.GetPublicAndNonAdultNonAdVideos(ctx)
	if err != nil {
		return nil, err
	}

	tags, err := i.db.Database.GetAllVideosTags(ctx)
	if err != nil {
		return nil, err
	}

	for _, dbVideo := range dbVideos {
		video := domain.NewVideo(dbVideo.ID, dbVideo.VideoUrl, dbVideo.ThumbnailImageUrl, dbVideo.Title, &dbVideo.Description.String, dbVideo.UploaderID, dbVideo.CreatedAt)
		for _, tag := range tags {
			if tag.VideoID == dbVideo.ID {
				video.Tags = append(video.Tags, tag.TagName)
			}
		}
		video.CreatedAt = time.Now()
		video.UpdatedAt = time.Now()
		videos = append(videos, video)
	}

	return videos, nil
}

func (i *Infrastructure) GetVideosByUserIDFromDB(ctx context.Context, userID string) ([]*domain.Video, error) {
	var videos []*domain.Video
	dbVideos, err := i.db.Database.GetPublicAndNonAdByUploaderID(ctx, userID)
	if err != nil {
		return nil, err
	}

	tags, err := i.db.Database.GetAllVideosTagsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, dbVideo := range dbVideos {
		video := domain.NewVideo(dbVideo.ID, dbVideo.VideoUrl, dbVideo.ThumbnailImageUrl, dbVideo.Title, &dbVideo.Description.String, dbVideo.UploaderID, dbVideo.CreatedAt)
		for _, tag := range tags {
			if tag.VideoID == dbVideo.ID {
				video.Tags = append(video.Tags, tag.TagName)
			}
			videos = append(videos, video)
		}
	}
	return videos, nil
}

func (i *Infrastructure) GetVideoFromDB(ctx context.Context, id string) (*domain.Video, error) {
	dbVideo, err := i.db.Database.GetVideo(ctx, id)
	if err != nil {
		return nil, err
	}

	tags, err := i.db.Database.GetVideoTags(ctx, id)
	if err != nil {
		return nil, err
	}

	video := domain.NewVideo(dbVideo.ID, dbVideo.VideoUrl, dbVideo.ThumbnailImageUrl, dbVideo.Title, &dbVideo.Description.String, dbVideo.UploaderID, time.Now()) // dbVideo.CreatedAt)
	for _, tag := range tags {
		video.Tags = append(video.Tags, tag.TagName)
	}

	return video, nil
}

func (i *Infrastructure) InsertVideo(ctx context.Context, id string, videoURL string, thumbnailImageURL string, title string, description *string, uploaderID string, tags []string, isAdult bool, isPrivate bool, isExternalCutout bool, isAd bool) (*domain.UploadVideoResponse, error) {
	_, err := i.db.Database.CreateVideo(ctx, sqlc.CreateVideoParams{
		ID:                id,
		VideoUrl:          videoURL,
		ThumbnailImageUrl: thumbnailImageURL,
		Title:             title,
		Description: sql.NullString{
			String: *description,
			Valid:  description != nil,
		},
		UploaderID:       uploaderID,
		IsPrivate:        isPrivate,
		IsAdult:          isAdult,
		IsExternalCutout: isExternalCutout,
		IsAd:             isAd,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		WatchCount:       0,
	})
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		tagInsertResult, err := i.db.Database.CreateTags(ctx, tag)
		if err != nil {
			return nil, err
		}

		tagID, err := tagInsertResult.LastInsertId()
		if err != nil {
			return nil, err
		}

		_, err = i.db.Database.CreateVideoTags(ctx, sqlc.CreateVideoTagsParams{
			VideoID: id,
			TagID:   int32(tagID),
		})
		if err != nil {
			return nil, err
		}
	}

	return &domain.UploadVideoResponse{
		ID:                id,
		VideoURL:          videoURL,
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Description:       description,
		UploaderID:        uploaderID,
		Tags:              tags,
		IsAdult:           isAdult,
		IsPrivate:         isPrivate,
		IsExternalCutout:  isExternalCutout,
		IsAd:              isAd,
		// CreatedAt:         time.Now(),
	}, nil
}
