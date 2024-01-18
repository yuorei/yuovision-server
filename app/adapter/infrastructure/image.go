package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/kolesa-team/go-webp/webp"
)

func (i *Infrastructure) ConvertThumbnailToWebp(ctx context.Context, imageFile *io.ReadSeeker, contentType string) (*bytes.Buffer, error) {
	if imageFile == nil {
		return nil, fmt.Errorf("there is no image")
	}

	var image image.Image
	switch contentType {
	case "image/jpeg":
		// JPEG画像をデコード
		img, err := jpeg.Decode(*imageFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode JPEG image")
		}
		image = img

	case "image/png":
		// PNG画像をデコード
		img, err := png.Decode(*imageFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode PNG image")
		}
		image = img
	case "image/webp":
		// WEBP画像をデコード
		img, err := webp.Decode(*imageFile, nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode WEBP image")
		}
		image = img
	default:
		return nil, fmt.Errorf("This file is not supported.")
	}

	// WebPに変換
	webpBuffer := new(bytes.Buffer)
	err := webp.Encode(webpBuffer, image, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode to WEBP image")
	}

	return webpBuffer, nil
}
