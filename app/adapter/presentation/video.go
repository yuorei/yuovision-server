package presentation

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
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

func (s *VideoService) Videos(ctx context.Context, _ *empty.Empty) (*video_grpc.VideosResponse, error) {
	videos, err := s.usecase.GetVideos(ctx)
	if err != nil {
		return nil, err
	}

	var videoPayloads []*video_grpc.VideoPayload
	for _, video := range videos {
		var description string
		if video.Description == nil {
			description = ""
		} else {
			description = *video.Description
		}

		videoPayloads = append(videoPayloads, &video_grpc.VideoPayload{
			Id:                video.ID,
			VideoUrl:          video.VideoURL,
			Title:             video.Title,
			ThumbnailImageUrl: video.ThumbnailImageURL,
			Description:       description,
			CreatedAt:         timestamppb.New(video.CreatedAt),
			UpdatedAt:         timestamppb.New(video.UpdatedAt),
			UserId:            video.UploaderID,
			Tags:              video.Tags,
			Private:           video.IsPrivate,
			Adult:             video.IsAdult,
			ExternalCutout:    video.IsExternalCutout,
			IsAd:              video.IsAd,
		})
	}

	return &video_grpc.VideosResponse{
		Videos: videoPayloads,
	}, nil
}

func (s *VideoService) VideosByUserID(ctx context.Context, id *video_grpc.VideoUserID) (*video_grpc.VideosResponse, error) {
	videos, err := s.usecase.GetVideosByUserID(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	var videoPayloads []*video_grpc.VideoPayload
	for _, video := range videos {
		videoPayloads = append(videoPayloads, &video_grpc.VideoPayload{
			Id:                video.ID,
			VideoUrl:          video.VideoURL,
			Title:             video.Title,
			ThumbnailImageUrl: video.ThumbnailImageURL,
			Description:       *video.Description,
			CreatedAt:         timestamppb.New(video.CreatedAt),
			UpdatedAt:         timestamppb.New(video.UpdatedAt),
			UserId:            video.UploaderID,
			Tags:              video.Tags,
			Private:           video.IsPrivate,
			Adult:             video.IsAdult,
			ExternalCutout:    video.IsExternalCutout,
			IsAd:              video.IsAd,
		})
	}

	return &video_grpc.VideosResponse{
		Videos: videoPayloads,
	}, nil
}

func (s *VideoService) UploadThumbnail(stream video_grpc.VideoService_UploadThumbnailServer) error {
	ctx := context.Background()
	var imageFile *os.File
	var id, contentType string

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
			case *video_grpc.UploadThumbnailInput_Meta:
				id = x.Meta.Id
				contentType = x.Meta.ContentType
				if x.Meta.ContentType != "" {
					contentType = strings.Split(contentType, "/")[1]
					imageFile, err = os.Create(id + "." + contentType)
					defer func() {
						if _, err := os.Stat(id + "." + contentType); err == nil {
							if err := os.Remove(id + "." + contentType); err != nil {
								return
							}
						}
					}()
					if err != nil {
						return err
					}
					defer imageFile.Close()
				}
			}
		}
	}
	// TODO
	// 既定サイズにきりとりする処理

	thumbnail := domain.NewThumbnailImage(id, contentType)
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
				err := os.MkdirAll(tempDir, 0755)
				if err != nil {
					return err
				}
				tempMp4 := filepath.Join(tempDir, id+".mp4")
				videoFile, err = os.Create(tempMp4)
				if err != nil {
					return err
				}
				defer videoFile.Close()
				defer func() {
					if _, err := os.Stat(tempMp4); err == nil {
						if err := os.Remove(tempMp4); err != nil {
							return
						}
					}
				}()
			}
		}
	}

	video := domain.NewUploadVideo(id, videoFile, meta.Title, &meta.Description, meta.Tags, meta.Adult, meta.Private, meta.ExternalCutout, meta.IsAd)
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
			Tags:              uploadVideo.Tags,
			Private:           uploadVideo.IsPrivate,
			Adult:             uploadVideo.IsAdult,
			ExternalCutout:    uploadVideo.IsExternalCutout,
			IsAd:              uploadVideo.IsAd,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *VideoService) WatchCount(ctx context.Context, id *video_grpc.WatchCountInput) (*video_grpc.WatchCountPayload, error) {
	watchCount, err := s.usecase.GetWatchCount(ctx, id.VideoId)
	if err != nil {
		return nil, err
	}

	return &video_grpc.WatchCountPayload{
		Count: int32(watchCount),
	}, nil
}

func (s *VideoService) IncrementWatchCount(ctx context.Context, input *video_grpc.IncrementWatchCountInput) (*video_grpc.WatchCountPayload, error) {
	watchCount, err := s.usecase.IncrementWatchCount(ctx, input.VideoId, input.UserId)
	if err != nil {
		return nil, err
	}

	return &video_grpc.WatchCountPayload{
		Count: int32(watchCount),
	}, nil
}

func (s *VideoService) CutVideo(ctx context.Context, input *video_grpc.CutVideoInput) (*video_grpc.CutVideoPayload, error) {
	url, err := s.usecase.CutVideo(ctx, input.VideoId, input.UserId, int(input.Start), int(input.End))
	if err != nil {
		return nil, err
	}

	return &video_grpc.CutVideoPayload{
		VideoUrl: url,
	}, nil
}
