package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/yuorei/video-server/directive"
	"github.com/yuorei/video-server/graph/generated"
	"github.com/yuorei/video-server/graph/resolver"
	"github.com/yuorei/video-server/middleware"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Could not load: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	c := generated.Config{Resolvers: &resolver.Resolver{}}
	c.Directives.Auth = directive.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("URL")},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"*"},
	})
	router.PathPrefix("/graphql").Handler(corsOpts.Handler(srv))
	router.PathPrefix("/").Handler(playground.Handler("GraphQL playground", "/graphql"))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func init() {
	yuorei :=
		`
██╗   ██╗██╗   ██╗ ██████╗ ██████╗ ███████╗██╗
╚██╗ ██╔╝██║   ██║██╔═══██╗██╔══██╗██╔════╝██║
 ╚████╔╝ ██║   ██║██║   ██║██████╔╝█████╗  ██║
  ╚██╔╝  ██║   ██║██║   ██║██╔══██╗██╔══╝  ██║
   ██║   ╚██████╔╝╚██████╔╝██║  ██║███████╗██║
   ╚═╝    ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝
`
	fmt.Println(yuorei)
}
