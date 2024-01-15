package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type CommentInputPort interface {
	PostComment(context.Context, *domain.Comment) (*domain.Comment, error)
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type CommentRepository interface {
	InsertComment(context.Context, *domain.Comment) (*domain.Comment, error)
}
