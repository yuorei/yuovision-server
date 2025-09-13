package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"cloud.google.com/go/firestore"
	firestoreRepo "github.com/yuorei/video-server/app/adapter/infrastructure/firestore"
	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/driver/firebase"
	"github.com/yuorei/video-server/app/driver/pubsub"
	"github.com/yuorei/video-server/app/driver/storage"
)

type Infrastructure struct {
	Video           port.VideoPort
	VideoProcessing port.VideoProcessingRepository
	User            port.UserPort
	Comment         port.CommentPort
	Image           port.ImagePort

	FirestoreClient *firestore.Client
	R2Client        *storage.R2Client
	FirebaseAuth    *firebase.AuthClient
	PubSubClient    *pubsub.Client
}

type R2Config = storage.R2Config

type InfraConfig struct {
	FirebaseCredentialsPath string
	FirebaseProjectID       string
	R2Config                R2Config
	PubSubConfig            PubSubConfig
}

type PubSubConfig struct {
	ProjectID       string
	CredentialsPath string
}

func NewInfrastructure(ctx context.Context, cfg InfraConfig) (*Infrastructure, error) {
	// Initialize Firestore
	firestoreClient, err := firebase.NewFirestoreClient(ctx, cfg.FirebaseProjectID, cfg.FirebaseCredentialsPath)
	if err != nil {
		return nil, err
	}

	// Initialize Firebase Auth
	authClient, err := firebase.NewAuthClient(cfg.FirebaseProjectID, cfg.FirebaseCredentialsPath)
	if err != nil {
		return nil, err
	}

	// Initialize R2 Storage
	r2Client, err := storage.NewR2Client(ctx, cfg.R2Config)
	if err != nil {
		return nil, err
	}

	// Initialize Pub/Sub Client
	var pubsubClient *pubsub.Client
	if cfg.PubSubConfig.ProjectID != "" && cfg.PubSubConfig.CredentialsPath != "" {
		client, err := pubsub.NewClient(ctx, pubsub.Config{
			ProjectID:       cfg.PubSubConfig.ProjectID,
			CredentialsPath: cfg.PubSubConfig.CredentialsPath,
		})
		if err != nil {
			slog.Error("failed to initialize Pub/Sub client", "project_id", cfg.PubSubConfig.ProjectID, "error", err)
			return nil, fmt.Errorf("failed to initialize Pub/Sub client with project ID '%s': %w", cfg.PubSubConfig.ProjectID, err)
		}
		pubsubClient = client
		slog.Info("Pub/Sub client initialized successfully", "project_id", cfg.PubSubConfig.ProjectID)
	} else {
		if cfg.PubSubConfig.ProjectID == "" {
			slog.Warn("GOOGLE_CLOUD_PROJECT_ID is empty, skipping Pub/Sub client initialization")
		}
		if cfg.PubSubConfig.CredentialsPath == "" {
			slog.Warn("GOOGLE_APPLICATION_CREDENTIALS is empty, skipping Pub/Sub client initialization")
		}
		pubsubClient = nil
	}

	// Initialize repositories
	videoRepo := firestoreRepo.NewVideoRepository(firestoreClient.Client())
	videoProcessingRepo := firestoreRepo.NewVideoProcessingRepository(firestoreClient.Client())
	userRepo := firestoreRepo.NewUserRepository(firestoreClient.Client())
	commentRepo := firestoreRepo.NewCommentRepository(firestoreClient.Client())

	// Initialize image service using R2
	imageRepo := NewImageRepository(r2Client)

	return &Infrastructure{
		Video:           videoRepo,
		VideoProcessing: videoProcessingRepo,
		User:            userRepo,
		Comment:         commentRepo,
		Image:           imageRepo,
		FirestoreClient: firestoreClient.Client(),
		R2Client:        r2Client,
		FirebaseAuth:    authClient,
		PubSubClient:    pubsubClient,
	}, nil
}

// ValidationVideo validates if the provided video is valid
func (i *Infrastructure) ValidationVideo(video io.ReadSeeker) error {
	if video == nil {
		return errors.New("video is nil")
	}

	// Read first few bytes to check file format
	buffer := make([]byte, 512)
	_, err := video.Read(buffer)
	if err != nil {
		return errors.New("failed to read video")
	}

	// Reset the reader position
	video.Seek(0, io.SeekStart)

	// Check for empty file or invalid format
	// This is a simplified validation - you might want to use a proper video format detection library
	if len(buffer) == 0 {
		return errors.New("video is empty")
	}

	// Simple format check based on file signatures
	if buffer[0] == 0x89 && buffer[1] == 0x50 && buffer[2] == 0x4E && buffer[3] == 0x47 {
		return errors.New("file is PNG, not a video")
	}

	return nil
}
