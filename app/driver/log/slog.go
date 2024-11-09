package log

import (
	"log/slog"
	"os"
)

func NewLog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
