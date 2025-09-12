package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type AuthClient struct {
	client *auth.Client
}

func NewAuthClient(credentialsPath string) (*AuthClient, error) {
	var app *firebase.App
	var err error

	if credentialsPath != "" {
		// Use credentials file if provided
		opt := option.WithCredentialsFile(credentialsPath)
		app, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing app with credentials file %s: %v", credentialsPath, err)
		}
	} else {
		// Use Application Default Credentials (ADC) for Cloud Run
		app, err = firebase.NewApp(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("error initializing app with ADC: %v", err)
		}
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return &AuthClient{client: client}, nil
}

func (ac *AuthClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := ac.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}

func (ac *AuthClient) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	u, err := ac.client.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user %s: %v", uid, err)
	}
	return u, nil
}
