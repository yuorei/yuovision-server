package storage

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Client struct {
	client *s3.Client
	bucket string
}

type R2Config struct {
	AccessKeyID     string
	SecretAccessKey string
	AccountID       string
	BucketName      string
}

func NewR2Client(ctx context.Context, cfg R2Config) (*R2Client, error) {
	slog.Info("initializing R2 client", "accountID", cfg.AccountID, "bucket", cfg.BucketName)

	if cfg.AccessKeyID == "" || cfg.SecretAccessKey == "" || cfg.AccountID == "" || cfg.BucketName == "" {
		return nil, fmt.Errorf("R2 configuration is incomplete")
	}

	// Cloudflare R2 endpoint
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.AccountID)
	slog.Info("using R2 endpoint", "endpoint", endpoint)

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	slog.Info("loading AWS config for R2")
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKeyID,
			cfg.SecretAccessKey,
			"",
		)),
		config.WithRegion("auto"), // R2 uses 'auto' as region
	)
	if err != nil {
		slog.Error("failed to load AWS config for R2", "error", err)
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}
	slog.Info("AWS config loaded successfully")

	client := s3.NewFromConfig(awsCfg)
	slog.Info("R2 client initialized successfully")

	return &R2Client{
		client: client,
		bucket: cfg.BucketName,
	}, nil
}

func (r2 *R2Client) UploadFile(ctx context.Context, key string, body io.Reader, contentType string) error {
	_, err := r2.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file %s: %w", key, err)
	}

	return nil
}

func (r2 *R2Client) GetPresignedURL(ctx context.Context, key string, expiration time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(r2.client)

	request, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r2.bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expiration
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL for %s: %w", key, err)
	}

	return request.URL, nil
}

func (r2 *R2Client) GetUploadPresignedURL(ctx context.Context, key string, contentType string, expiration time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(r2.client)

	request, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expiration
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate upload presigned URL for %s: %w", key, err)
	}

	return request.URL, nil
}

func (r2 *R2Client) DeleteFile(ctx context.Context, key string) error {
	_, err := r2.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r2.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", key, err)
	}

	return nil
}
