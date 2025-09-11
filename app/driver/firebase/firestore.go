package firebase

import (
	"context"
	"fmt"
	"log/slog"

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
	slog.Info("initializing Firestore client", "projectID", projectID, "credentialsPath", credentialsPath != "")

	var app *firebase.App
	var err error

	if credentialsPath != "" {
		// Use credentials file if provided
		slog.Info("using credentials file for Firestore")
		opt := option.WithCredentialsFile(credentialsPath)
		app, err = firebase.NewApp(ctx, &firebase.Config{
			ProjectID: projectID,
		}, opt)
	} else {
		// Use Application Default Credentials (ADC) for Cloud Run
		slog.Info("using Application Default Credentials for Firestore")
		app, err = firebase.NewApp(ctx, &firebase.Config{
			ProjectID: projectID,
		})
	}

	if err != nil {
		slog.Error("failed to initialize Firebase app", "error", err)
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}
	slog.Info("Firebase app initialized successfully")

	client, err := app.Firestore(ctx)
	if err != nil {
		slog.Error("failed to initialize Firestore client", "error", err)
		return nil, fmt.Errorf("error initializing firestore client: %v", err)
	}
	slog.Info("Firestore client initialized successfully")

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
