package presentation

import (
	"context"

	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func NewUserService(app *application.Application) *UserService {
	return &UserService{
		usecase: application.NewUseCase(app),
	}
}

type UserService struct {
	video_grpc.UnimplementedUserServiceServer
	usecase *application.UseCase
}

func (s *UserService) User(ctx context.Context, input *video_grpc.UserID) (*video_grpc.UserPayload, error) {
	user, err := s.usecase.UserInputPort.GetUser(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	return &video_grpc.UserPayload{
		Id:                  user.ID,
		Name:                user.Name,
		ProfileImageUrl:     user.ProfileImageURL,
		SubscribeChannelIds: user.Subscribechannelids,
		IsSubscribed:        user.IsSubscribed,
		// Role:                 user.Role,
	}, nil
}

func (s *UserService) RegisterUser(ctx context.Context, input *video_grpc.UserInput) (*video_grpc.UserPayload, error) {
	user := domain.NewUser(input.Id, input.Name, input.ProfileImageUrl, input.SubscribeChannelIds, input.IsSubscribed, input.Role.String())
	user, err := s.usecase.UserInputPort.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &video_grpc.UserPayload{
		Id:                  user.ID,
		Name:                user.Name,
		ProfileImageUrl:     user.ProfileImageURL,
		SubscribeChannelIds: user.Subscribechannelids,
		IsSubscribed:        user.IsSubscribed,
		// Role:                user.Role,
	}, nil
}

func (s *UserService) SubscribeChannel(ctx context.Context, input *video_grpc.SubscribeChannelInput) (*video_grpc.SubscriptionPayload, error) {
	subscribeChannel := domain.NewSubscribeChannel(input.UserId, input.ChannelId)
	subscribeChannel, err := s.usecase.UserInputPort.SubscribeChannel(ctx, subscribeChannel)
	if err != nil {
		return nil, err
	}

	return &video_grpc.SubscriptionPayload{
		IsSuccess: subscribeChannel.IsSuccess,
	}, nil
}

func (s *UserService) UnSubscribeChannel(ctx context.Context, input *video_grpc.SubscribeChannelInput) (*video_grpc.SubscriptionPayload, error) {
	subscribeChannel := domain.NewSubscribeChannel(input.UserId, input.ChannelId)
	subscribeChannel, err := s.usecase.UserInputPort.UnSubscribeChannel(ctx, subscribeChannel)
	if err != nil {
		return nil, err
	}

	return &video_grpc.SubscriptionPayload{
		IsSuccess: subscribeChannel.IsSuccess,
	}, nil
}
