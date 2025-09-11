package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type FirestoreClient struct {
	client *firestore.Client
}

func (fc *FirestoreClient) Client() *firestore.Client {
	return fc.client
}

func NewFirestoreClient(ctx context.Context, projectID, credentialsPath string) (*FirestoreClient, error) {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: projectID,
	}, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore client: %v", err)
	}

	return &FirestoreClient{client: client}, nil
}

func (fc *FirestoreClient) Close() error {
	return fc.client.Close()
}

func (fc *FirestoreClient) Collection(name string) *firestore.CollectionRef {
	return fc.client.Collection(name)
}

func (fc *FirestoreClient) Doc(path string) *firestore.DocumentRef {
	return fc.client.Doc(path)
}
