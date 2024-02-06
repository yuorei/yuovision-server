package collection

type (
	User struct {
		ID                  string `bson:"_id"`
		Name                string
		ProfileImageURL     string
		SubscribeChannelIDs []string
	}
)

func NewUserCollection(id, name, profileImageURL string) *User {
	return &User{
		ID:              id,
		Name:            name,
		ProfileImageURL: profileImageURL,
	}
}
