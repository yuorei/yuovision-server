package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (i *Infrastructure) ConvertVideoHLS(ctx context.Context, videoID string) error {
	// HLS変換の実行
	outputDir := "output/" + videoID
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	output := "output_" + videoID + ".m3u8"
	outputHLS := filepath.Join(outputDir, output)
	tempMp4 := filepath.Join("temp", videoID+".mp4")
	cmd := exec.Command("ffmpeg", "-i", tempMp4, "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", outputHLS, "-y")
	log.Println(cmd.Args)
	result, err := cmd.CombinedOutput()
	log.Println(string(result))
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	return nil
}
