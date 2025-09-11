package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type CommentUseCase struct {
	commentRepo port.CommentPort
}

func NewCommentUseCase(commentRepo port.CommentPort) *CommentUseCase {
	return &CommentUseCase{
		commentRepo: commentRepo,
	}
}

func (uc *CommentUseCase) GetCommentsByVideoID(ctx context.Context, videoID string) ([]*domain.Comment, error) {
	return uc.commentRepo.GetByVideoID(ctx, videoID)
}

func (uc *CommentUseCase) GetComment(ctx context.Context, commentID string) (*domain.Comment, error) {
	return uc.commentRepo.GetByID(ctx, commentID)
}

func (uc *CommentUseCase) CreateComment(ctx context.Context, comment *domain.Comment) error {
	return uc.commentRepo.Create(ctx, comment)
}
