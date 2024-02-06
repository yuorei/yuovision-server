package application

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/lib"
	mock_port "github.com/yuorei/video-server/mock"
)

func TestGetVideos_SortsByCreatedAtDescending(t *testing.T) {
	type testCase struct {
		name        string
		mockVideos  []*domain.Video
		wantVideos  []*domain.Video
		expectError bool
	}

	tests := []testCase{
		{
			name: "Sort videos by created date in descending order",
			mockVideos: []*domain.Video{
				{
					ID:                "video1",
					VideoURL:          "https://example.com/video1",
					ThumbnailImageURL: "https://example.com/thumbnail1",
					Title:             "video1",
					Description:       lib.StringPointer("description1"),
					UploaderID:        "user1",
					CreatedAt:         time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:                "video2",
					VideoURL:          "https://example.com/video2",
					ThumbnailImageURL: "https://example.com/thumbnail2",
					Title:             "video2",
					Description:       lib.StringPointer("description2"),
					UploaderID:        "user2",
					CreatedAt:         time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:                "video3",
					VideoURL:          "https://example.com/video3",
					ThumbnailImageURL: "https://example.com/thumbnail3",
					Title:             "video3",
					Description:       nil,
					UploaderID:        "user3",
					CreatedAt:         time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			wantVideos: []*domain.Video{
				{
					ID:                "video2",
					VideoURL:          "https://example.com/video2",
					ThumbnailImageURL: "https://example.com/thumbnail2",
					Title:             "video2",
					Description:       lib.StringPointer("description2"),
					UploaderID:        "user2",
					CreatedAt:         time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:                "video3",
					VideoURL:          "https://example.com/video3",
					ThumbnailImageURL: "https://example.com/thumbnail3",
					Title:             "video3",
					Description:       nil,
					UploaderID:        "user3",
					CreatedAt:         time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:                "video1",
					VideoURL:          "https://example.com/video1",
					ThumbnailImageURL: "https://example.com/thumbnail1",
					Title:             "video1",
					Description:       lib.StringPointer("description1"),
					UploaderID:        "user1",
					CreatedAt:         time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockVideoRepository := mock_port.NewMockVideoRepository(controller)
			mockVideoRepository.EXPECT().GetVideosFromDB(ctx).Return(tc.mockVideos, nil)

			app := NewApplication(nil)
			videoRepositor := NewVideoUseCase(mockVideoRepository)
			app.Video = videoRepositor
			usecase := &UseCase{
				VideoInputPort: app,
			}

			// Execute the test subject method
			gotVideos, err := usecase.GetVideos(ctx)
			if (err != nil) != tc.expectError {
				t.Fatalf("unexpected error: %v, expectError: %v", err, tc.expectError)
			}

			// Compare the result
			if diff := cmp.Diff(gotVideos, tc.wantVideos); diff != "" {
				t.Errorf("Response body mismatch: %s", diff)
			}
		})
	}
}
