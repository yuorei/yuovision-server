package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/app/driver/db/mongodb/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	user := domain.NewUser(userForDB.ID, userForDB.Name, userForDB.ProfileImageURL, userForDB.SubscribeChannelIDs)
	return user, nil
}

func (i *Infrastructure) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	mongoCollection := i.db.Database.Collection("user")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	userForDB := collection.NewUserCollection(user.ID, user.Name, user.ProfileImageURL)
	insertResult, err := mongoCollection.InsertOne(ctx, userForDB)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return user, nil
}

func (i *Infrastructure) GetProfileImageURL(ctx context.Context, id string) (string, error) {
	resp, err := http.Get(os.Getenv("AUTH_URL") + "/profile-image/" + id)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスの処理
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var profileImageURL domain.ProfileImageURL
	err = json.Unmarshal(body, &profileImageURL)
	if err != nil {
		return "", err
	}

	return profileImageURL.URL, nil
}

func (i *Infrastructure) AddSubscribeChannelForDB(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	mongoCollection := i.db.Database.Collection("user")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	// チャンネル登録していないかを確認している
	var result bson.M
	err := mongoCollection.FindOne(ctx, bson.M{"subscribechannelids": subscribeChannel.ChannelID}).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("already subscribed")
	}

	//チャンネルが存在するかを確認する
	err = mongoCollection.FindOne(ctx, bson.M{"_id": subscribeChannel.ChannelID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Channel does not exist")
		}
		return nil, fmt.Errorf("error while checking ChannelID existence: %v", err)
	}

	// チャンネルを登録する
	filter := bson.M{"_id": subscribeChannel.UserID}
	update := bson.M{
		"$addToSet": bson.M{"subscribechannelids": subscribeChannel.ChannelID},
	}
	options := options.Update().SetUpsert(true)

	_, err = mongoCollection.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return nil, fmt.Errorf("error while updating user: %v", err)
	}
	subscribeChannel.IsSuccess = true

	return subscribeChannel, nil
}

func (i *Infrastructure) UnSubscribeChannelForDB(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	mongoCollection := i.db.Database.Collection("user")
	if mongoCollection == nil {
		return nil, fmt.Errorf("collection is nil")
	}

	// チャンネル登録しているかを確認している
	var result bson.M
	err := mongoCollection.FindOne(ctx, bson.M{"subscribechannelids": subscribeChannel.ChannelID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("ChannelID does not exist")
		}
		return nil, err
	}
	//チャンネルが存在するかを確認する
	err = mongoCollection.FindOne(ctx, bson.M{"_id": subscribeChannel.ChannelID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Channel does not exist")
		}
		return nil, fmt.Errorf("error while checking ChannelID existence: %v", err)
	}

	// チャンネルを解除する
	filter := bson.M{"_id": subscribeChannel.UserID}
	update := bson.M{
		"$pull": bson.M{"subscribechannelids": subscribeChannel.ChannelID},
	}

	_, err = mongoCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error while updating user: %v", err)
	}
	subscribeChannel.IsSuccess = true

	return subscribeChannel, nil
}
