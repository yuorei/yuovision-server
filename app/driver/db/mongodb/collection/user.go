package collection

type (
	User struct {
		ID   string `bson:"_id"`
		Name string
	}
)

func NewUserCollection(id string, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
