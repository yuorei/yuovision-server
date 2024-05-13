package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) ConvertVideoHLS(ctx context.Context, video *domain.VideoFile) error {
	// HLS変換の実行
	outputDir := "output"
	os.MkdirAll(outputDir, 0755)

	output := "output_" + video.ID + ".m3u8"
	outputHLS := filepath.Join(outputDir, output)
	tempMp4 := filepath.Join("temp", video.ID+".mp4")
	cmd := exec.Command("ffmpeg", "-i", tempMp4, "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", outputHLS, "-y")
	log.Println(cmd.Args)
	result, err := cmd.CombinedOutput()
	log.Println(string(result))
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	return nil
}
