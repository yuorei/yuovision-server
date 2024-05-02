package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
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

func (a *Application) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return a.User.userRepository.InsertUser(ctx, user)
}

func (a *Application) SubscribeChannel(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	return a.User.userRepository.AddSubscribeChannelForDB(ctx, subscribeChannel)
}

func (a *Application) UnSubscribeChannel(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	return a.User.userRepository.UnSubscribeChannelForDB(ctx, subscribeChannel)
}
