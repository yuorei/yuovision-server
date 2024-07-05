package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/db/sqlc"
)

func (i *Infrastructure) GetUserFromDB(ctx context.Context, id string) (*domain.User, error) {
	user, err := i.db.Database.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	subscribechannelids, err := i.db.Database.GetUserSubscribeChannelsID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:                  user.ID,
		Name:                user.Name,
		ProfileImageURL:     user.ProfileImageUrl,
		Subscribechannelids: subscribechannelids,
	}, nil
}

func (i *Infrastructure) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	result, err := i.db.Database.CreatetUser(ctx, sqlc.CreatetUserParams{
		ID:              user.ID,
		Name:            user.Name,
		ProfileImageUrl: user.ProfileImageURL,
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("failed to insert user")
	}
	log.Println("inserted user", user.ID)
	return user, nil

}

func (i *Infrastructure) GetProfileImageURL(ctx context.Context, id string) (string, error) {
	resp, err := http.Get(os.Getenv("AUTH_URL") + "/profile-image/" + id)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスの処理
	body, err := io.ReadAll(resp.Body)
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
	id, err := i.db.Database.GetUserSubscriptionID(ctx, sqlc.GetUserSubscriptionIDParams{
		UserID:    subscribeChannel.UserID,
		ChannelID: subscribeChannel.ChannelID,
	})

	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
	}
	if id != "" {
		return nil, fmt.Errorf("already subscribed")
	}

	_, err = i.db.Database.SubscribeChannel(ctx, sqlc.SubscribeChannelParams{
		UserID:    subscribeChannel.UserID,
		ChannelID: subscribeChannel.ChannelID,
	})
	if err != nil {
		return nil, err
	}

	subscribeChannel.IsSuccess = true
	return subscribeChannel, nil
}

func (i *Infrastructure) UnSubscribeChannelForDB(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	id, err := i.db.Database.GetUserSubscriptionID(ctx, sqlc.GetUserSubscriptionIDParams{
		UserID:    subscribeChannel.UserID,
		ChannelID: subscribeChannel.ChannelID,
	})
	if err != nil {
		return nil, err
	}
	if id == "" {
		return nil, fmt.Errorf("not subscribed")
	}

	_, err = i.db.Database.UnSubscribeChannel(ctx, sqlc.UnSubscribeChannelParams{
		UserID:    subscribeChannel.UserID,
		ChannelID: subscribeChannel.ChannelID,
	})
	if err != nil {
		return nil, err
	}

	subscribeChannel.IsSuccess = true
	return subscribeChannel, nil
}
