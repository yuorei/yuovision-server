package port

import (
	"context"
	"io"

	"github.com/yuorei/video-server/app/domain"
)

// VideoPort - Firestore video repository interface
type VideoPort interface {
	Create(context.Context, *domain.Video) error
	GetByID(context.Context, string) (*domain.Video, error)
	GetAll(context.Context) ([]*domain.Video, error)
	GetVideosByUserID(context.Context, string) ([]*domain.Video, error)
	Update(context.Context, *domain.Video) error
	Delete(context.Context, string) error
}

// UserPort - Firestore user repository interface
type UserPort interface {
	Create(context.Context, *domain.User) error
	GetByID(context.Context, string) (*domain.User, error)
	GetAll(context.Context) ([]*domain.User, error)
	Update(context.Context, *domain.User) error
}

// CommentPort - Firestore comment repository interface
type CommentPort interface {
	Create(context.Context, *domain.Comment) error
	GetByID(context.Context, string) (*domain.Comment, error)
	GetByVideoID(context.Context, string) ([]*domain.Comment, error)
}

// ImagePort - R2 storage interface
type ImagePort interface {
	Upload(context.Context, string, io.Reader, string) (string, error)
	GetPresignedURL(context.Context, string) (string, error)
	GetUploadURL(context.Context, string, string) (string, error)
	Delete(context.Context, string) error
}
