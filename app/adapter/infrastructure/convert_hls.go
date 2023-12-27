package infrastructure

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) ConvertVideoHLS(ctx context.Context, video *domain.VideoFile) error {
	tempDir := "temp"
	os.MkdirAll(tempDir, 0755)
	tempMp4 := filepath.Join(tempDir, video.ID+".mp4")

	// 一時ファイルの作成
	tempFile, err := os.Create(tempMp4)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// アップロードされたファイルの内容を一時ファイルにコピー
	_, err = io.Copy(tempFile, video.Video)
	if err != nil {
		return err
	}

	// HLS変換の実行
	outputDir := "output"
	os.MkdirAll(outputDir, 0755)

	output := "output_" + video.ID + ".m3u8"
	outputHLS := filepath.Join(outputDir, output)
	cmd := exec.Command("ffmpeg", "-i", tempMp4, "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", outputHLS)
	log.Println(cmd.Args)
	result, err := cmd.CombinedOutput()
	log.Println(string(result))
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	// 一時ファイルの削除
	os.Remove(tempMp4)

	return nil
}
