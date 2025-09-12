package pubsub

import (
	"context"
	"fmt"
	"log/slog"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type Client struct {
	client *pubsub.Client
}

type Config struct {
	ProjectID       string
	CredentialsPath string
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	slog.Info("initializing Pub/Sub client", "project_id", cfg.ProjectID, "credentials_path", cfg.CredentialsPath)
	
	if cfg.ProjectID == "" {
		return nil, fmt.Errorf("pubsub: projectID string is empty")
	}
	
	var client *pubsub.Client
	var err error

	if cfg.CredentialsPath != "" {
		slog.Info("using credentials file for Pub/Sub")
		client, err = pubsub.NewClient(ctx, cfg.ProjectID, option.WithCredentialsFile(cfg.CredentialsPath))
	} else {
		slog.Info("using Application Default Credentials for Pub/Sub")
		client, err = pubsub.NewClient(ctx, cfg.ProjectID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create pub/sub client: %w", err)
	}

	slog.Info("Pub/Sub client initialized successfully")
	return &Client{client: client}, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) PublishVideoProcessingMessage(ctx context.Context, topicID string, data []byte) error {
	topic := c.client.Topic(topicID)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (c *Client) SubscribeVideoProcessing(ctx context.Context, subscriptionID string, handler func(context.Context, []byte) error) error {
	sub := c.client.Subscription(subscriptionID)

	return sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if err := handler(ctx, msg.Data); err != nil {
			msg.Nack()
			return
		}
		msg.Ack()
	})
}
