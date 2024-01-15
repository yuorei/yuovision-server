package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type CommentUseCase struct {
	commentRepository port.CommentRepository
}

func NewCommentUseCase(commentRepository port.CommentRepository) *CommentUseCase {
	return &CommentUseCase{
		commentRepository: commentRepository,
	}
}

func (a *Application) PostComment(ctx context.Context, postComment *domain.Comment) (*domain.Comment, error) {
	return a.Comment.commentRepository.InsertComment(ctx, postComment)
}
