package firestore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/yuorei/video-server/app/domain"
)

type VideoRepository struct {
	client     *firestore.Client
	collection string
}

func NewVideoRepository(client *firestore.Client) *VideoRepository {
	return &VideoRepository{
		client:     client,
		collection: "videos",
	}
}

type VideoDoc struct {
	ID                string    `firestore:"id"`
	VideoURL          string    `firestore:"video_url"`
	ThumbnailImageURL string    `firestore:"thumbnail_image_url"`
	Title             string    `firestore:"title"`
	Description       *string   `firestore:"description"`
	Tags              []string  `firestore:"tags"`
	WatchCount        int       `firestore:"watch_count"`
	IsPrivate         bool      `firestore:"is_private"`
	IsAdult           bool      `firestore:"is_adult"`
	IsExternalCutout  bool      `firestore:"is_external_cutout"`
	IsAd              bool      `firestore:"is_ad"`
	UploaderID        string    `firestore:"uploader_id"`
	CreatedAt         time.Time `firestore:"created_at"`
	UpdatedAt         time.Time `firestore:"updated_at"`
}

func (r *VideoRepository) Create(ctx context.Context, video *domain.Video) error {
	doc := VideoDoc{
		ID:                video.ID,
		VideoURL:          video.VideoURL,
		ThumbnailImageURL: video.ThumbnailImageURL,
		Title:             video.Title,
		Description:       video.Description,
		Tags:              video.Tags,
		WatchCount:        video.WatchCount,
		IsPrivate:         video.IsPrivate,
		IsAdult:           video.IsAdult,
		IsExternalCutout:  video.IsExternalCutout,
		IsAd:              video.IsAd,
		UploaderID:        video.UploaderID,
		CreatedAt:         video.CreatedAt,
		UpdatedAt:         video.UpdatedAt,
	}

	_, err := r.client.Collection(r.collection).Doc(video.ID).Set(ctx, doc)
	if err != nil {
		return fmt.Errorf("failed to create video %s: %w", video.ID, err)
	}

	return nil
}

func (r *VideoRepository) GetByID(ctx context.Context, id string) (*domain.Video, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get video %s: %w", id, err)
	}

	var videoDoc VideoDoc
	if err := doc.DataTo(&videoDoc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal video data: %w", err)
	}

	video := &domain.Video{
		ID:                videoDoc.ID,
		VideoURL:          videoDoc.VideoURL,
		ThumbnailImageURL: videoDoc.ThumbnailImageURL,
		Title:             videoDoc.Title,
		Description:       videoDoc.Description,
		Tags:              videoDoc.Tags,
		WatchCount:        videoDoc.WatchCount,
		IsPrivate:         videoDoc.IsPrivate,
		IsAdult:           videoDoc.IsAdult,
		IsExternalCutout:  videoDoc.IsExternalCutout,
		IsAd:              videoDoc.IsAd,
		UploaderID:        videoDoc.UploaderID,
		CreatedAt:         videoDoc.CreatedAt,
		UpdatedAt:         videoDoc.UpdatedAt,
	}

	return video, nil
}

func (r *VideoRepository) GetAll(ctx context.Context) ([]*domain.Video, error) {
	iter := r.client.Collection(r.collection).
		Where("is_private", "==", false).
		OrderBy("created_at", firestore.Desc).
		Documents(ctx)

	var videos []*domain.Video
	for {
		doc, err := iter.Next()
		if err != nil {
			if err.Error() == "iterator stopped" {
				break
			}
			return nil, fmt.Errorf("failed to iterate videos: %w", err)
		}

		var videoDoc VideoDoc
		if err := doc.DataTo(&videoDoc); err != nil {
			continue // Skip invalid documents
		}

		video := &domain.Video{
			ID:                videoDoc.ID,
			VideoURL:          videoDoc.VideoURL,
			ThumbnailImageURL: videoDoc.ThumbnailImageURL,
			Title:             videoDoc.Title,
			Description:       videoDoc.Description,
			Tags:              videoDoc.Tags,
			WatchCount:        videoDoc.WatchCount,
			IsPrivate:         videoDoc.IsPrivate,
			IsAdult:           videoDoc.IsAdult,
			IsExternalCutout:  videoDoc.IsExternalCutout,
			IsAd:              videoDoc.IsAd,
			UploaderID:        videoDoc.UploaderID,
			CreatedAt:         videoDoc.CreatedAt,
			UpdatedAt:         videoDoc.UpdatedAt,
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (r *VideoRepository) Update(ctx context.Context, video *domain.Video) error {
	doc := VideoDoc{
		ID:                video.ID,
		VideoURL:          video.VideoURL,
		ThumbnailImageURL: video.ThumbnailImageURL,
		Title:             video.Title,
		Description:       video.Description,
		Tags:              video.Tags,
		WatchCount:        video.WatchCount,
		IsPrivate:         video.IsPrivate,
		IsAdult:           video.IsAdult,
		IsExternalCutout:  video.IsExternalCutout,
		IsAd:              video.IsAd,
		UploaderID:        video.UploaderID,
		CreatedAt:         video.CreatedAt,
		UpdatedAt:         time.Now(),
	}

	_, err := r.client.Collection(r.collection).Doc(video.ID).Set(ctx, doc, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("failed to update video %s: %w", video.ID, err)
	}

	return nil
}

func (r *VideoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection(r.collection).Doc(id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete video %s: %w", id, err)
	}

	return nil
}
