package router

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/yuorei/video-server/app/adapter/infrastructure"
	"github.com/yuorei/video-server/app/adapter/presentation/resolver"
	"github.com/yuorei/video-server/app/application"
	flog "github.com/yuorei/video-server/app/driver/log"
	"github.com/yuorei/video-server/app/driver/sentry"
	"github.com/yuorei/video-server/graph/generated"
)

const (
	defaultHTTPPort = "8080"
)

// maskSecret masks secret values for logging
func maskSecret(secret string) string {
	if secret == "" {
		return ""
	}
	if len(secret) <= 4 {
		return "****"
	}
	return secret[:2] + "****" + secret[len(secret)-2:]
}

func NewHTTPRouter() {
	flog.NewLog()
	slog.Info("start HTTP GraphQL server")

	sentry.SentryInit()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultHTTPPort
	}
	slog.Info("using port", "port", port)

	// Log environment variables for debugging
	slog.Info("environment variables",
		"FIREBASE_CREDENTIALS_PATH", os.Getenv("FIREBASE_CREDENTIALS_PATH"),
		"FIREBASE_PROJECT_ID", os.Getenv("FIREBASE_PROJECT_ID"),
		"R2_ACCESS_KEY_ID", maskSecret(os.Getenv("R2_ACCESS_KEY_ID")),
		"R2_SECRET_ACCESS_KEY", maskSecret(os.Getenv("R2_SECRET_ACCESS_KEY")),
		"R2_ACCOUNT_ID", os.Getenv("R2_ACCOUNT_ID"),
		"R2_BUCKET_NAME", os.Getenv("R2_BUCKET_NAME"),
	)

	// Initialize new infrastructure with Firebase and R2
	infraConfig := infrastructure.InfraConfig{
		FirebaseCredentialsPath: os.Getenv("FIREBASE_CREDENTIALS_PATH"),
		FirebaseProjectID:       os.Getenv("FIREBASE_PROJECT_ID"),
		R2Config: infrastructure.R2Config{
			AccessKeyID:     os.Getenv("R2_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("R2_SECRET_ACCESS_KEY"),
			AccountID:       os.Getenv("R2_ACCOUNT_ID"),
			BucketName:      os.Getenv("R2_BUCKET_NAME"),
		},
	}

	slog.Info("initializing infrastructure")
	infra, err := infrastructure.NewInfrastructure(context.Background(), infraConfig)
	if err != nil {
		slog.Error("failed to initialize infrastructure", "error", err)
		log.Fatalf("Failed to initialize infrastructure: %v", err)
	}
	slog.Info("infrastructure initialized successfully")

	app := application.NewApplication(infra)
	resolver := resolver.NewResolver(app)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Info("starting GraphQL server", "port", port, "playground", "http://localhost:"+port+"/")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	slog.Info("server listening on port", "port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("server failed to start", "error", err)
		log.Fatal(err)
	}
}
