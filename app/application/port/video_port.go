package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type VideoInputPort interface {
	GetVideo(context.Context, string) (*domain.Video, error)
	UploadVideo(context.Context, *domain.UploadVideo) (*domain.UploadVideoResponse, error)
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type VideoRepository interface {
	ConvertVideoHLS(context.Context, *domain.VideoFile) error
	UploadVideoForStorage(context.Context, *domain.VideoFile) (string, error)
	GetVideoFromDB(context.Context, string) (*domain.Video, error)
	InsertVideo(context.Context, string, string, string, string, *string, string) (*domain.UploadVideoResponse, error)
}
