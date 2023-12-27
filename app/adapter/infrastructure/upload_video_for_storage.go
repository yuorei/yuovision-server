package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) UploadVideoForStorage(ctx context.Context, video *domain.VideoFile) (*domain.UploadVideoForStorageResponse, error) {
	err := filepath.Walk("output", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 対象のファイルかどうかを確認
		if strings.HasPrefix(filepath.Base(path), "output_"+video.ID) && (strings.HasSuffix(path, ".m3u8") || strings.HasSuffix(path, ".ts")) {
			// TODO: 失敗した時にtsファイルを削除できるように修正する
			defer func() error {
				err = os.Remove(path)
				if err != nil {
					return err
				}
				return nil
			}()
			err := uploadForS3(path)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to remove output files: %w", err)
	}

	return &domain.UploadVideoForStorageResponse{}, nil
}

func uploadForS3(path string) error {
	ctx := context.Background()
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	endpoint := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: "http://localhost:9000",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred), config.WithEndpointResolver(endpoint))
	if err != nil {
		return err
	}

	// change object address style
	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})

	// get buckets
	lbo, err := client.ListBuckets(ctx, nil)
	if err != nil {
		return err
	}
	buckets := make(map[string]struct{}, len(lbo.Buckets))
	for _, b := range lbo.Buckets {
		buckets[*b.Name] = struct{}{}
	}

	// create 'video-service' bucket if not exist
	bucketName := "video-service"
	if _, ok := buckets[bucketName]; !ok {
		_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: &bucketName,
		})
		if err != nil {
			return err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// put object
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &strings.Split(path, "/")[1],
		Body:   file,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	log.Println("Successful upload: ", path)

	return nil
}
