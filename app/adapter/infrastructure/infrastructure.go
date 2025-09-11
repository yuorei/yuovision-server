package infrastructure

import (
	"context"
	"errors"
	"io"

	"cloud.google.com/go/firestore"
	firestoreRepo "github.com/yuorei/video-server/app/adapter/infrastructure/firestore"
	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/driver/firebase"
	"github.com/yuorei/video-server/app/driver/storage"
)

type Infrastructure struct {
	Video   port.VideoPort
	User    port.UserPort
	Comment port.CommentPort
	Image   port.ImagePort

	FirestoreClient *firestore.Client
	R2Client        *storage.R2Client
	FirebaseAuth    *firebase.AuthClient
}

type R2Config = storage.R2Config

type InfraConfig struct {
	FirebaseCredentialsPath string
	FirebaseProjectID       string
	R2Config                R2Config
}

func NewInfrastructure(ctx context.Context, cfg InfraConfig) (*Infrastructure, error) {
	// Initialize Firestore
	firestoreClient, err := firebase.NewFirestoreClient(ctx, cfg.FirebaseProjectID, cfg.FirebaseCredentialsPath)
	if err != nil {
		return nil, err
	}

	// Initialize Firebase Auth
	authClient, err := firebase.NewAuthClient(cfg.FirebaseCredentialsPath)
	if err != nil {
		return nil, err
	}

	// Initialize R2 Storage
	r2Client, err := storage.NewR2Client(ctx, cfg.R2Config)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	videoRepo := firestoreRepo.NewVideoRepository(firestoreClient.Client())
	userRepo := firestoreRepo.NewUserRepository(firestoreClient.Client())
	commentRepo := firestoreRepo.NewCommentRepository(firestoreClient.Client())

	// Initialize image service using R2
	imageRepo := NewImageRepository(r2Client)

	return &Infrastructure{
		Video:           videoRepo,
		User:            userRepo,
		Comment:         commentRepo,
		Image:           imageRepo,
		FirestoreClient: firestoreClient.Client(),
		R2Client:        r2Client,
		FirebaseAuth:    authClient,
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
