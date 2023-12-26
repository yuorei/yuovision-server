package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) UploadVideoForStorage(video *domain.VideoFile) (*domain.UploadVideoForStorageResponse, error) {
	// TODO ストレージサービスに保存する
	err := filepath.Walk("output", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 対象のファイルかどうかを確認
		if strings.HasPrefix(filepath.Base(path), "output_"+video.ID) && (strings.HasSuffix(path, ".m3u8") || strings.HasSuffix(path, ".ts")) {
			err := os.Remove(path)
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
