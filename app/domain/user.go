package domain

type (
	User struct {
		ID              string
		Name            string
		ProfileImageURL string
	}

	SubscribeChannel struct {
		UserID    string
		ChannelID string
		IsSuccess bool
	}

	ProfileImageURL struct {
		URL string `json:"url"`
	}
)

func NewUser(id, name, profileImageURL string) *User {
	return &User{
		ID:              id,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}
}

func NewSubscribeChannel(userID, channelID string) *SubscribeChannel {
	return &SubscribeChannel{
		UserID:    userID,
		ChannelID: channelID,
		IsSuccess: false,
	}
}
