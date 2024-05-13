package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type UserInputPort interface {
	GetUser(context.Context, string) (*domain.User, error)
	RegisterUser(context.Context, *domain.User) (*domain.User, error)
	SubscribeChannel(context.Context, *domain.SubscribeChannel) (*domain.SubscribeChannel, error)
	UnSubscribeChannel(context.Context, *domain.SubscribeChannel) (*domain.SubscribeChannel, error)
}

type UserRepository interface {
	GetProfileImageURL(context.Context, string) (string, error)
	GetUserFromDB(context.Context, string) (*domain.User, error)
	InsertUser(context.Context, *domain.User) (*domain.User, error)
	AddSubscribeChannelForDB(context.Context, *domain.SubscribeChannel) (*domain.SubscribeChannel, error)
	UnSubscribeChannelForDB(context.Context, *domain.SubscribeChannel) (*domain.SubscribeChannel, error)
}
