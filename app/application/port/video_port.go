package port

import (
	"context"
	"io"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type VideoInputPort interface {
	GetVideos(context.Context) ([]*domain.Video, error)
	GetVideosByUserID(context.Context, string) ([]*domain.Video, error)
	GetVideo(context.Context, string) (*domain.Video, error)
	UploadVideo(context.Context, *domain.UploadVideo, string, string) (*domain.UploadVideoResponse, error)
	GetWatchCount(context.Context, string) (int, error)
	IncrementWatchCount(context.Context, string, string) (int, error)
	CutVideo(context.Context, string, string, int, int) (string, error)
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type VideoRepository interface {
	CheckUploadAPIRateLimit(context.Context, string) error
	SetUploadAPIRateLimit(context.Context, string) error
	GetVideosFromDB(context.Context) ([]*domain.Video, error)
	GetVideosByUserIDFromDB(context.Context, string) ([]*domain.Video, error)
	ConvertVideoHLS(context.Context, string) error
	ValidationVideo(io.ReadSeeker) error
	UploadVideoForStorage(context.Context, *domain.VideoFile) (string, error)
	GetVideoFromDB(context.Context, string) (*domain.Video, error)
	InsertVideo(context.Context, string, string, string, string, *string, string, []string, bool, bool, bool, bool) (*domain.UploadVideoResponse, error)
	GetWatchCount(context.Context, string) (int, error)
	ChechWatchCount(context.Context, string, string) (bool, error)
	IncrementWatchCount(context.Context, string, string) (int, error)
	CutVideo(context.Context, string, string, int, int) (string, error)
}
