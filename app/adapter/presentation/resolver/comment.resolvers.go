package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
	model "github.com/yuorei/video-server/app/domain/models"
	"github.com/yuorei/video-server/middleware"
)

// PostComment is the resolver for the PostComment field.
func (r *mutationResolver) PostComment(ctx context.Context, input model.PostCommentInput) (*model.PostCommentPayload, error) {
	commentID := domain.NewCommentID()
	userID, err := middleware.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	name, err := middleware.GetNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	postComment := domain.NewPostComment(commentID, input.VideoID, userID, name, input.Text)
	postComment, err = r.usecase.PostComment(ctx, postComment)
	if err != nil {
		return nil, err
	}

	return &model.PostCommentPayload{
		ID:        postComment.ID,
		VideoID:   postComment.VideoID,
		Text:      postComment.Text,
		CreatedAt: postComment.CreatedAt.String(),
		UpdatedAt: postComment.UpdatedAt.String(),
		User: &model.User{
			ID:   postComment.User.ID,
			Name: postComment.User.Name,
		},
	}, nil
}
