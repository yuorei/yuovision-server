package router

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yuorei/video-server/graph/generated"

	"github.com/yuorei/video-server/app/adapter/infrastructure"
	resolver "github.com/yuorei/video-server/app/adapter/presentation/resolver"
	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/middleware"
)

func NewRouter() {
	const defaultPort = "8080"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	infra := infrastructure.NewInfrastructure()
	app := application.NewApplication(infra)
	r := resolver.NewResolver(app)
	c := generated.Config{Resolvers: r}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// Logger:         log.New(os.Stdout, "video-server", log.LstdFlags),
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"*"},
	})

	router := mux.NewRouter()
	router.Use(middleware.Middleware())
	router.PathPrefix("/graphql").Handler(corsOpts.Handler(srv))
	router.PathPrefix("/").Handler(playground.Handler("GraphQL playground", "/graphql"))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
