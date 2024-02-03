package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/middleware"
)

type UserUseCase struct {
	userRepository port.UserRepository
}

func NewUserUseCase(userRepository port.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (a *Application) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return a.User.userRepository.GetUserFromDB(ctx, id)
}

func (a *Application) RegisterUser(ctx context.Context) (*domain.User, error) {
	id, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	name, err := middleware.GetNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 画像URLを取得
	profileImageURL, err := a.User.userRepository.GetProfileImageURL(ctx, id)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(id, name, profileImageURL)
	return a.User.userRepository.InsertUser(ctx, user)
}
