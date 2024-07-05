// ImageRepository
package port

import (
	"context"
	"os"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type ImageInputPort interface {
	UploadThumbnail(context.Context, domain.ThumbnailImage) error
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type ImageRepository interface {
	ConvertThumbnailToWebp(context.Context, *os.File, string, string) (*os.File, error)
	UploadImageForStorage(context.Context, string) (string, error)
	CreateThumbnail(context.Context, string) error
}
