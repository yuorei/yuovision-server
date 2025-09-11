package sentry

import (
	"log/slog"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func SentryInit() {
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		slog.Warn("Sentry DSN is empty, skipping Sentry initialization")
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		slog.Error("failed to initialize Sentry", "error", err)
		return
	}

	slog.Info("Sentry initialized successfully")
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("It works!")
}
