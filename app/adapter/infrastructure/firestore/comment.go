package firestore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"github.com/yuorei/video-server/app/domain"
)

type CommentRepository struct {
	client     *firestore.Client
	collection string
}

func NewCommentRepository(client *firestore.Client) *CommentRepository {
	return &CommentRepository{
		client:     client,
		collection: "comments",
	}
}

type CommentDoc struct {
	ID        string    `firestore:"id"`
	VideoID   string    `firestore:"video_id"`
	UserID    string    `firestore:"user_id"`
	Text      string    `firestore:"text"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

func (r *CommentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	doc := CommentDoc{
		ID:        comment.ID,
		VideoID:   comment.VideoID,
		UserID:    comment.UserID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	_, err := r.client.Collection(r.collection).Doc(comment.ID).Set(ctx, doc)
	if err != nil {
		return fmt.Errorf("failed to create comment %s: %w", comment.ID, err)
	}

	return nil
}

func (r *CommentRepository) GetByID(ctx context.Context, id string) (*domain.Comment, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment %s: %w", id, err)
	}

	var commentDoc CommentDoc
	if err := doc.DataTo(&commentDoc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal comment data: %w", err)
	}

	comment := &domain.Comment{
		ID:        commentDoc.ID,
		VideoID:   commentDoc.VideoID,
		UserID:    commentDoc.UserID,
		Text:      commentDoc.Text,
		CreatedAt: commentDoc.CreatedAt,
		UpdatedAt: commentDoc.UpdatedAt,
	}

	return comment, nil
}

func (r *CommentRepository) GetByVideoID(ctx context.Context, videoID string) ([]*domain.Comment, error) {
	iter := r.client.Collection(r.collection).
		Where("video_id", "==", videoID).
		OrderBy("created_at", firestore.Desc).
		Documents(ctx)

	var comments []*domain.Comment
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("failed to iterate comments: %w", err)
		}

		var commentDoc CommentDoc
		if err := doc.DataTo(&commentDoc); err != nil {
			continue // Skip invalid documents
		}

		comment := &domain.Comment{
			ID:        commentDoc.ID,
			VideoID:   commentDoc.VideoID,
			UserID:    commentDoc.UserID,
			Text:      commentDoc.Text,
			CreatedAt: commentDoc.CreatedAt,
			UpdatedAt: commentDoc.UpdatedAt,
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
