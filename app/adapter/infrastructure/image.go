package infrastructure

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/kolesa-team/go-webp/webp"
)

func (i *Infrastructure) ConvertThumbnailToWebp(ctx context.Context, imageFile *os.File, contentType, id string) (*os.File, error) {
	if imageFile == nil {
		return nil, nil
	}

	var image image.Image
	switch "image/" + contentType {
	case "image/jpeg":
		// JPEG画像をデコード
		img, err := jpeg.Decode(imageFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode JPEG image")
		}
		image = img

	case "image/png":
		// PNG画像をデコード
		img, err := png.Decode(imageFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode PNG image")
		}
		image = img

	case "image/webp":
		return imageFile, nil
	default:
		return nil, fmt.Errorf("This file is not supported: %s", contentType)
	}

	imageTmp, err := os.Create(id + ".webp")
	if err != nil {
		return nil, err
	}
	defer imageTmp.Close()
	// WebPにエンコード
	err = webp.Encode(imageTmp, image, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode to WEBP image")
	}

	return imageTmp, nil
}

func (i *Infrastructure) UploadImageForStorage(ctx context.Context, id string) (string, error) {
	imagePath := id + ".webp"
	defer func() error {
		err := os.Remove(imagePath)
		if err != nil {
			return err
		}
		return nil
	}()
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred))
	if err != nil {
		return "", err
	}

	// change object address style
	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
		options.BaseEndpoint = aws.String(os.Getenv("AWS_S3_ENDPOINT"))
		options.Region = "ap-northeast-1"
	})

	// get buckets
	lbo, err := client.ListBuckets(ctx, nil)
	if err != nil {
		return "", err
	}
	buckets := make(map[string]struct{}, len(lbo.Buckets))
	for _, b := range lbo.Buckets {
		buckets[*b.Name] = struct{}{}
	}

	// create 'thumbnail-image' bucket if not exist
	bucketName := "thumbnail-image"
	if _, ok := buckets[bucketName]; !ok {
		_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: &bucketName,
			ACL:    types.BucketCannedACLPublicRead,
		})
		if err != nil {
			return "", err
		}
	}

	image, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer image.Close()

	// bytes.bufferがaws-sdk-go-v2では使えない
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imagePath),
		Body:   image,
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}
	log.Println("Successful upload: ", imagePath)

	url := fmt.Sprintf("%s/%s/%s.webp", os.Getenv("AWS_S3_URL"), bucketName, id)
	return url, nil
}

func (i *Infrastructure) CreateThumbnail(ctx context.Context, id string) error {
	const bucketName = "video"
	imagePath := id + ".webp"
	tmpVideoPath := id + ".mp4"
	defer func() error {
		err := os.Remove(tmpVideoPath)
		if err != nil {
			return err
		}
		return nil
	}()

	// S3から取得HLS
	// 処理を2回に分けているのはこの方法が早いため
	// 分けないとffmpegが全てのファイルをダウンロードしてから処理を行うため時間がかかってしまう
	url := fmt.Sprintf("%s/%s/output_%s.m3u8", os.Getenv("AWS_S3_URL"), bucketName, id)
	cmd := exec.Command("ffmpeg", "-ss", "00:00:00", "-t", "1", "-i", url, tmpVideoPath)
	log.Println(cmd.Args)
	result, err := cmd.CombinedOutput()
	log.Println(string(result))
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	cmd = exec.Command("ffmpeg", "-i", tmpVideoPath, "-vframes", "1", imagePath)
	log.Println(cmd.Args)
	result, err = cmd.CombinedOutput()
	log.Println(string(result))
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	return nil
}
