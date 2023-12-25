package port

import "github.com/yuorei/video-server/app/domain"

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type VideoInputPort interface {
	UploadVideo(video *domain.UploadVideo) (*domain.UploadVideoResponse, error)
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type VideoRepository interface {
	ConvertVideoHLS(video *domain.VideoFile) error
	UploadVideoForStorage(video *domain.VideoFile) (*domain.UploadVideoForStorageResponse, error)
}
