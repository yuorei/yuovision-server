package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	collection := i.db.Database.Collection("user")
	if collection == nil {
		return nil, fmt.Errorf("collection is nil")
	}
	insertResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return user, nil
}
