package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type UserUseCase struct {
	userRepo port.UserPort
}

func NewUserUseCase(userRepo port.UserPort) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return uc.userRepo.GetAll(ctx)
}

func (uc *UserUseCase) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	return uc.userRepo.GetByID(ctx, userID)
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	return uc.userRepo.Create(ctx, user)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	return uc.userRepo.Update(ctx, user)
}
