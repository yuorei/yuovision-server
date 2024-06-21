package router

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/yuorei/video-server/app/adapter/infrastructure"
	"github.com/yuorei/video-server/app/adapter/presentation"
	"github.com/yuorei/video-server/app/application"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func NewRouter() {
	const defaultPort = "50051"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	infra := infrastructure.NewInfrastructure()
	app := application.NewApplication(infra)

	video_grpc.RegisterUserServiceServer(s, presentation.NewUserService(app))
	video_grpc.RegisterCommentServiceServer(s, presentation.NewCommentService(app))
	video_grpc.RegisterVideoServiceServer(s, presentation.NewVideoService(app))
	reflection.Register(s)
	log.Printf("start gRPC server port: %v", port)
	s.Serve(listener)
}
