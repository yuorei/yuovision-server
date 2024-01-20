// ImageRepository
package port

import (
	"bytes"
	"context"
	"io"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type ImageInputPort interface {
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type ImageRepository interface {
	ConvertThumbnailToWebp(context.Context, *io.ReadSeeker, string) (*bytes.Buffer, error)
	UploadImageForStorage(context.Context, string, *bytes.Buffer) (string, error)
}
