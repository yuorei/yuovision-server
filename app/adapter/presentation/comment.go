package presentation

import (
	"context"

	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCommentService(app *application.Application) *CommentService {
	return &CommentService{
		usecase: application.NewUseCase(app),
	}
}

type CommentService struct {
	video_grpc.UnimplementedCommentServiceServer
	usecase *application.UseCase
}

func (s *CommentService) CommentsByVideo(ctx context.Context, id *video_grpc.CommentsByVideoInput) (*video_grpc.CommentsResponse, error) {
	comments, err := s.usecase.GetCommentsByVideoID(ctx, id.VideoId)
	if err != nil {
		return nil, err
	}

	var commentPayloads []*video_grpc.Comment
	for _, comment := range comments {
		commentPayloads = append(commentPayloads, &video_grpc.Comment{
			Id:     comment.ID,
			UserId: comment.User.ID,
			Text:   comment.Text,
			Name:   comment.User.Name,
			// IsOwner:   isOwner,
			CreatedAt: timestamppb.New(comment.CreatedAt),
			UpdatedAt: timestamppb.New(comment.UpdatedAt),
			Video: &video_grpc.Video{
				Id: comment.VideoID,
			},
		})
	}

	return &video_grpc.CommentsResponse{
		Comments: commentPayloads,
	}, nil
}

func (s *CommentService) PostComment(ctx context.Context, input *video_grpc.PostCommentInput) (*video_grpc.PostCommentPayload, error) {
	commentID := domain.NewCommentID()
	comment := domain.NewPostComment(commentID, input.VideoId, input.UserId, input.Name, input.Text)
	comment, err := s.usecase.PostComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &video_grpc.PostCommentPayload{
		Id:     comment.ID,
		UserId: comment.User.ID,
		Text:   comment.Text,
		Name:   comment.User.Name,
		// IsOwner:   comment.IsOwner,
		CreatedAt: timestamppb.New(comment.CreatedAt),
		UpdatedAt: timestamppb.New(comment.UpdatedAt),
		Video: &video_grpc.Video{
			Id: comment.VideoID,
		},
	}, nil
}
