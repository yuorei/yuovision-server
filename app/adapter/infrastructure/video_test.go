package infrastructure

import (
	"io"
	"os"
	"testing"
)

func Test_動画のバリデーションチェック(t *testing.T) {
	type fields struct {
		// db    *db.DB
		// redis *redis.Client
	}
	type args struct {
		video io.ReadSeeker
	}

	video1, err := os.Open("test1.mp4")
	if err != nil {
		t.Errorf("failed to open video file")
	}
	videoMOV, err := os.Open("test.MOV")
	if err != nil {
		t.Errorf("failed to open video file")
	}
	video2, err := os.Open("test2.mp4")
	if err != nil {
		t.Errorf("failed to open video file")
	}
	video3, err := os.Open("test3.mp4")
	if err != nil {
		t.Errorf("failed to open video file")
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "success mp4",
			fields: fields{},
			args: args{
				video: video1,
			},
			wantErr: false,
		},
		{
			name:   "success MOV",
			fields: fields{},
			args: args{
				video: videoMOV,
			},
			wantErr: false,
		},
		{
			name:   "video is nil",
			fields: fields{},
			args: args{
				video: nil,
			},
			wantErr: true,
		},
		{
			name:   "video is empty",
			fields: fields{},
			args: args{
				video: video2,
			},
			wantErr: true,
		},
		{
			name:   "video is png",
			fields: fields{},
			args: args{
				video: video3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Infrastructure{}
			if err := i.ValidationVideo(tt.args.video); (err != nil) != tt.wantErr {
				t.Errorf("Infrastructure.ValidationVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
