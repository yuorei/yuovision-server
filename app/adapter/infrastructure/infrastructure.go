package infrastructure

import (
	"context"

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
