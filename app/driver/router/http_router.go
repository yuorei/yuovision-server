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

func NewHTTPRouter() {
	flog.NewLog()
	slog.Info("start HTTP GraphQL server")

	sentry.SentryInit()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultHTTPPort
	}

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

	infra, err := infrastructure.NewInfrastructure(context.Background(), infraConfig)
	if err != nil {
		log.Fatalf("Failed to initialize infrastructure: %v", err)
	}

	app := application.NewApplication(infra)
	resolver := resolver.NewResolver(app)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
