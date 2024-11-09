package sentry

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func SentryInit() {
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		log.Fatal("Error: Sentry DSN is empty")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	fmt.Println("Sentry Init")
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("It works!")
}
