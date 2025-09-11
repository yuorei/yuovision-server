package db

import (
	"context"
	"fmt"

	"github.com/yuorei/video-server/app/driver/firebase"
)

type DB struct {
	Firestore *firebase.FirestoreClient
}

func NewFirestoreDB(ctx context.Context, projectID, credentialsPath string) (*DB, error) {
	firestoreClient, err := firebase.NewFirestoreClient(ctx, projectID, credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create Firestore client: %w", err)
	}

	return &DB{
		Firestore: firestoreClient,
	}, nil
}

func (db *DB) Close() error {
	return db.Firestore.Close()
}
