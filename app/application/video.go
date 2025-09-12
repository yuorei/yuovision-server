package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type VideoUseCase struct {
	videoRepo           port.VideoPort
	videoProcessingRepo port.VideoProcessingRepository
	storageClient       StorageClient
	pubsubClient        PubSubClient
}

type StorageClient interface {
	UploadFile(ctx context.Context, key string, body io.Reader, contentType string) error
}

type PubSubClient interface {
	PublishVideoProcessingMessage(ctx context.Context, topicID string, data []byte) error
}

func NewVideoUseCase(videoRepo port.VideoPort) *VideoUseCase {
	return &VideoUseCase{
		videoRepo: videoRepo,
	}
}

func NewVideoUseCaseWithClients(
	videoRepo port.VideoPort,
	videoProcessingRepo port.VideoProcessingRepository,
	storageClient StorageClient,
	pubsubClient PubSubClient,
) *VideoUseCase {
	return &VideoUseCase{
		videoRepo:           videoRepo,
		videoProcessingRepo: videoProcessingRepo,
		storageClient:       storageClient,
		pubsubClient:        pubsubClient,
	}
}

func (uc *VideoUseCase) GetVideos(ctx context.Context) ([]*domain.Video, error) {
	videos, err := uc.videoRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[j].CreatedAt.Before(videos[i].CreatedAt)
	})

	return videos, nil
}

func (uc *VideoUseCase) GetVideo(ctx context.Context, videoID string) (*domain.Video, error) {
	return uc.videoRepo.GetByID(ctx, videoID)
}

func (uc *VideoUseCase) GetWatchCount(ctx context.Context, videoID string) (int, error) {
	video, err := uc.videoRepo.GetByID(ctx, videoID)
	if err != nil {
		return 0, err
	}
	return video.WatchCount, nil
}

func (uc *VideoUseCase) IncrementWatchCount(ctx context.Context, videoID, userID string) (int, error) {
	video, err := uc.videoRepo.GetByID(ctx, videoID)
	if err != nil {
		return 0, err
	}

	video.WatchCount++
	err = uc.videoRepo.Update(ctx, video)
	if err != nil {
		return 0, err
	}

	return video.WatchCount, nil
}

func (uc *VideoUseCase) CutVideo(ctx context.Context, videoID string, start, end int) (string, error) {
	// TODO: Implement video cutting logic
	return "", nil
}

func (uc *VideoUseCase) UploadThumbnailToStorage(ctx context.Context, key string, thumbnail io.ReadSeeker, contentType string) error {
	if uc.storageClient == nil {
		return fmt.Errorf("storage client not available")
	}
	return uc.storageClient.UploadFile(ctx, key, thumbnail, contentType)
}

func (uc *VideoUseCase) GetVideoProcessingInfo(ctx context.Context, videoID string) (*domain.VideoProcessingInfo, error) {
	if uc.videoProcessingRepo == nil {
		return nil, fmt.Errorf("video processing repository not available")
	}
	return uc.videoProcessingRepo.GetByVideoID(ctx, videoID)
}

func (uc *VideoUseCase) GetUserVideoProcessing(ctx context.Context, userID string) ([]*domain.VideoProcessingInfo, error) {
	if uc.videoProcessingRepo == nil {
		return nil, fmt.Errorf("video processing repository not available")
	}
	return uc.videoProcessingRepo.GetByUserID(ctx, userID)
}

func (uc *VideoUseCase) UploadVideo(ctx context.Context, uploadVideo *domain.UploadVideo, uploaderID string, thumbnailImageURL string) (*domain.UploadVideoResponse, error) {
	// Upload raw video file to storage
	videoKey := fmt.Sprintf("raw-videos/%s/original.mp4", uploadVideo.ID)
	if uc.storageClient != nil {
		err := uc.storageClient.UploadFile(ctx, videoKey, uploadVideo.Video, "video/mp4")
		if err != nil {
			return nil, fmt.Errorf("failed to upload video file: %w", err)
		}
	}

	// Create video entity with initial data
	video := domain.NewVideo(
		uploadVideo.ID,
		"", // VideoURL will be set after HLS processing
		thumbnailImageURL,
		uploadVideo.Title,
		uploadVideo.Description,
		uploadVideo.Tags,
		0, // Initial watch count
		uploadVideo.IsPrivate,
		uploadVideo.IsAdult,
		uploadVideo.IsExternalCutout,
		uploadVideo.IsAd,
		uploaderID,
		time.Now(),
		time.Now(),
	)

	// Save to repository
	err := uc.videoRepo.Create(ctx, video)
	if err != nil {
		return nil, fmt.Errorf("failed to create video: %w", err)
	}

	// Create initial processing info
	processingInfo := domain.NewVideoProcessingInfo(
		domain.NewVideoProcessingInfoID(),
		video.ID,
		domain.VideoProcessingStatusUploaded,
		0,
		nil,
		time.Now(),
		time.Now(),
	)

	if uc.videoProcessingRepo != nil {
		err = uc.videoProcessingRepo.Create(ctx, processingInfo)
		if err != nil {
			// Log error but don't fail the upload
			// TODO: Add proper logging
		}
	}

	// Send message to Pub/Sub for async processing
	if uc.pubsubClient != nil {
		message := VideoProcessingMessage{
			VideoID:          video.ID,
			VideoKey:         videoKey,
			ProcessingID:     processingInfo.ID,
			UploaderID:       uploaderID,
			Title:            video.Title,
			IsPrivate:        video.IsPrivate,
			IsAdult:          video.IsAdult,
			IsExternalCutout: video.IsExternalCutout,
		}

		messageData, err := json.Marshal(message)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal processing message: %w", err)
		}

		err = uc.pubsubClient.PublishVideoProcessingMessage(ctx, "video-processing", messageData)
		if err != nil {
			// Log error but don't fail the upload
			// TODO: Add proper logging
		}
	}

	// Return response
	return &domain.UploadVideoResponse{
		ID:                video.ID,
		VideoURL:          video.VideoURL,
		ThumbnailImageURL: video.ThumbnailImageURL,
		Title:             video.Title,
		Description:       video.Description,
		UploaderID:        video.UploaderID,
		Tags:              video.Tags,
		IsAdult:           video.IsAdult,
		IsPrivate:         video.IsPrivate,
		IsExternalCutout:  video.IsExternalCutout,
		IsAd:              video.IsAd,
		CreatedAt:         video.CreatedAt,
	}, nil
}

type VideoProcessingMessage struct {
	VideoID          string `json:"video_id"`
	VideoKey         string `json:"video_key"`
	ProcessingID     string `json:"processing_id"`
	UploaderID       string `json:"uploader_id"`
	Title            string `json:"title"`
	IsPrivate        bool   `json:"is_private"`
	IsAdult          bool   `json:"is_adult"`
	IsExternalCutout bool   `json:"is_external_cutout"`
}
