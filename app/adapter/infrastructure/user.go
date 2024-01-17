package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/app/driver/db/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *Infrastructure) GetUserFromDB(ctx context.Context, id string) (*domain.User, error) {
	mongoCollection := i.db.Database.Collection("user")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	var userForDB collection.User
	err := mongoCollection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&userForDB)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(userForDB.ID, userForDB.Name)
	return user, nil
}

func (i *Infrastructure) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	mongoCollection := i.db.Database.Collection("user")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	userForDB := collection.NewUserCollection(user.ID, user.Name)
	insertResult, err := mongoCollection.InsertOne(ctx, userForDB)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return user, nil
}
