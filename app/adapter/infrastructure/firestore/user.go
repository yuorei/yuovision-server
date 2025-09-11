package firestore

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/yuorei/video-server/app/domain"
	"google.golang.org/api/iterator"
)

type UserRepository struct {
	client     *firestore.Client
	collection string
}

func NewUserRepository(client *firestore.Client) *UserRepository {
	return &UserRepository{
		client:     client,
		collection: "users",
	}
}

type UserDoc struct {
	ID                  string    `firestore:"id"`
	Name                string    `firestore:"name"`
	ProfileImageURL     string    `firestore:"profile_image_url"`
	IsSubscribed        bool      `firestore:"is_subscribed"`
	Role                string    `firestore:"role"`
	Subscribechannelids []string  `firestore:"subscribe_channel_ids"`
	CreatedAt           time.Time `firestore:"created_at"`
	UpdatedAt           time.Time `firestore:"updated_at"`
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	doc := UserDoc{
		ID:                  user.ID,
		Name:                user.Name,
		ProfileImageURL:     user.ProfileImageURL,
		IsSubscribed:        user.IsSubscribed,
		Role:                user.Role,
		Subscribechannelids: user.Subscribechannelids,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	}

	_, err := r.client.Collection(r.collection).Doc(user.ID).Set(ctx, doc)
	if err != nil {
		return fmt.Errorf("failed to create user %s: %w", user.ID, err)
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %s: %w", id, err)
	}

	var userDoc UserDoc
	if err := doc.DataTo(&userDoc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	user := &domain.User{
		ID:                  userDoc.ID,
		Name:                userDoc.Name,
		ProfileImageURL:     userDoc.ProfileImageURL,
		IsSubscribed:        userDoc.IsSubscribed,
		Role:                userDoc.Role,
		Subscribechannelids: userDoc.Subscribechannelids,
		CreatedAt:           userDoc.CreatedAt,
		UpdatedAt:           userDoc.UpdatedAt,
	}

	return user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	iter := r.client.Collection(r.collection).Documents(ctx)

	var users []*domain.User
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("failed to iterate users: %w", err)
		}

		var userDoc UserDoc
		if err := doc.DataTo(&userDoc); err != nil {
			slog.Warn("Failed to unmarshal user document", "error", err, "document_id", doc.Ref.ID)
			continue // Skip invalid documents
		}

		user := &domain.User{
			ID:                  userDoc.ID,
			Name:                userDoc.Name,
			ProfileImageURL:     userDoc.ProfileImageURL,
			IsSubscribed:        userDoc.IsSubscribed,
			Role:                userDoc.Role,
			Subscribechannelids: userDoc.Subscribechannelids,
			CreatedAt:           userDoc.CreatedAt,
			UpdatedAt:           userDoc.UpdatedAt,
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	doc := UserDoc{
		ID:                  user.ID,
		Name:                user.Name,
		ProfileImageURL:     user.ProfileImageURL,
		IsSubscribed:        user.IsSubscribed,
		Role:                user.Role,
		Subscribechannelids: user.Subscribechannelids,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	}

	_, err := r.client.Collection(r.collection).Doc(user.ID).Set(ctx, doc, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("failed to update user %s: %w", user.ID, err)
	}

	return nil
}
