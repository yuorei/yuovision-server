package firestore

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/yuorei/video-server/app/domain"
	"google.golang.org/api/iterator"
)

type VideoProcessingRepository struct {
	client     *firestore.Client
	collection string
}

func NewVideoProcessingRepository(client *firestore.Client) *VideoProcessingRepository {
	return &VideoProcessingRepository{
		client:     client,
		collection: "video_processing",
	}
}

type VideoProcessingDoc struct {
	ID        string    `firestore:"id"`
	VideoID   string    `firestore:"video_id"`
	Status    string    `firestore:"status"`
	Progress  int       `firestore:"progress"`
	Message   *string   `firestore:"message"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

func (r *VideoProcessingRepository) Create(ctx context.Context, info *domain.VideoProcessingInfo) error {
	doc := VideoProcessingDoc{
		ID:        info.ID,
		VideoID:   info.VideoID,
		Status:    string(info.Status),
		Progress:  info.Progress,
		Message:   info.Message,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}

	_, err := r.client.Collection(r.collection).Doc(info.ID).Set(ctx, doc)
	if err != nil {
		return fmt.Errorf("failed to create video processing info %s: %w", info.ID, err)
	}

	return nil
}

func (r *VideoProcessingRepository) GetByVideoID(ctx context.Context, videoID string) (*domain.VideoProcessingInfo, error) {
	iter := r.client.Collection(r.collection).
		Where("video_id", "==", videoID).
		OrderBy("created_at", firestore.Desc).
		Limit(1).
		Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		if err == iterator.Done {
			return nil, fmt.Errorf("video processing info not found for video %s", videoID)
		}
		return nil, fmt.Errorf("failed to get video processing info for video %s: %w", videoID, err)
	}

	var processingDoc VideoProcessingDoc
	if err := doc.DataTo(&processingDoc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal video processing data: %w", err)
	}

	processingInfo := &domain.VideoProcessingInfo{
		ID:        processingDoc.ID,
		VideoID:   processingDoc.VideoID,
		Status:    domain.VideoProcessingStatus(processingDoc.Status),
		Progress:  processingDoc.Progress,
		Message:   processingDoc.Message,
		CreatedAt: processingDoc.CreatedAt,
		UpdatedAt: processingDoc.UpdatedAt,
	}

	return processingInfo, nil
}

func (r *VideoProcessingRepository) GetByUserID(ctx context.Context, userID string) ([]*domain.VideoProcessingInfo, error) {
	// First, get video IDs for this user
	videosIter := r.client.Collection("videos").
		Where("uploader_id", "==", userID).
		Documents(ctx)

	var videoIDs []string
	for {
		doc, err := videosIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("failed to iterate videos: %w", err)
		}
		videoIDs = append(videoIDs, doc.Ref.ID)
	}

	if len(videoIDs) == 0 {
		return []*domain.VideoProcessingInfo{}, nil
	}

	// Then get processing info for these videos
	var processingInfos []*domain.VideoProcessingInfo
	for _, videoID := range videoIDs {
		processingInfo, err := r.GetByVideoID(ctx, videoID)
		if err != nil {
			slog.Warn("Failed to get processing info for video", "video_id", videoID, "error", err)
			continue
		}
		processingInfos = append(processingInfos, processingInfo)
	}

	return processingInfos, nil
}

func (r *VideoProcessingRepository) Update(ctx context.Context, info *domain.VideoProcessingInfo) error {
	doc := VideoProcessingDoc{
		ID:        info.ID,
		VideoID:   info.VideoID,
		Status:    string(info.Status),
		Progress:  info.Progress,
		Message:   info.Message,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}

	_, err := r.client.Collection(r.collection).Doc(info.ID).Set(ctx, doc, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("failed to update video processing info %s: %w", info.ID, err)
	}

	return nil
}
