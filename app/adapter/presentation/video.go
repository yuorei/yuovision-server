package presentation

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewVideoService(app *application.Application) *VideoService {
	return &VideoService{
		usecase: application.NewUseCase(app),
	}
}

type VideoService struct {
	video_grpc.UnimplementedVideoServiceServer
	usecase *application.UseCase
}

func (s *VideoService) Video(ctx context.Context, id *video_grpc.VideoID) (*video_grpc.VideoPayload, error) {
	video, err := s.usecase.GetVideo(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	return &video_grpc.VideoPayload{
		Id:                video.ID,
		VideoUrl:          video.VideoURL,
		Title:             video.Title,
		ThumbnailImageUrl: video.ThumbnailImageURL,
		Description:       *video.Description,
		CreatedAt:         timestamppb.New(video.CreatedAt),
		UpdatedAt:         timestamppb.New(video.UpdatedAt),
		UserId:            video.UploaderID,
	}, nil
}

func (s *VideoService) UploadThumbnail(stream video_grpc.VideoService_UploadThumbnailServer) error {
	ctx := context.Background()
	var imageFile *os.File
	var id string

	for {
		input, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if input.GetValue() != nil {
			switch x := input.GetValue().(type) {
			case *video_grpc.UploadThumbnailInput_ThumbnailImage:
				_, err := imageFile.Write(x.ThumbnailImage)
				if err != nil {
					return err
				}
			case *video_grpc.UploadThumbnailInput_Id:
				id = x.Id
				imageFile, err = os.Create(id + ".webp")
				if err != nil {
					return err
				}
				defer imageFile.Close()
			}
		}
	}
	// TODO
	// サムネがなかった場合にサムネを作る処理
	// WEBPに変換する処理
	// 既定サイズにきりとりする処理
	// 画像をアップロードする処理 done

	thumbnail := domain.NewThumbnailImage(id)
	err := s.usecase.UploadThumbnail(ctx, thumbnail)
	if err != nil {
		return err
	}

	err = stream.SendAndClose(
		&video_grpc.UploadThumbnailPayload{
			Success: true,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *VideoService) UploadVideo(stream video_grpc.VideoService_UploadVideoServer) error {
	ctx := context.Background()
	var videoFile *os.File
	var id string
	var meta *video_grpc.VideoMeta

	for {
		input, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if input.GetValue() != nil {
			switch x := input.GetValue().(type) {
			case *video_grpc.UploadVideoInput_Video:
				_, err := videoFile.Write(x.Video)
				if err != nil {
					return err
				}
			case *video_grpc.UploadVideoInput_Meta:
				meta = x.Meta
				id = x.Meta.Id
				if id == "" {
					return fmt.Errorf("id is required")
				}
				tempDir := "temp"
				os.MkdirAll(tempDir, 0755)
				tempMp4 := filepath.Join(tempDir, id+".mp4")
				videoFile, err = os.Create(tempMp4)
				if err != nil {
					return err
				}
				defer videoFile.Close()
			}
		}
	}

	video := domain.NewUploadVideo(id, videoFile, meta.Title, &meta.Description)
	uploadVideo, err := s.usecase.UploadVideo(ctx, video, meta.UserId, meta.ThumbnailImageUrl)
	if err != nil {
		return err
	}

	err = stream.SendAndClose(
		&video_grpc.VideoPayload{
			Id:                uploadVideo.ID,
			VideoUrl:          uploadVideo.VideoURL,
			Title:             uploadVideo.Title,
			ThumbnailImageUrl: uploadVideo.ThumbnailImageURL,
			Description:       *uploadVideo.Description,
			CreatedAt:         timestamppb.New(uploadVideo.CreatedAt),
			UpdatedAt:         timestamppb.New(uploadVideo.CreatedAt),
			UserId:            uploadVideo.UploaderID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
