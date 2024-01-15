package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type UserInputPort interface {
	RegisterUser(context.Context) (*domain.User, error)
}

type UserRepository interface {
	InsertUser(context.Context, *domain.User) (*domain.User, error)
}
