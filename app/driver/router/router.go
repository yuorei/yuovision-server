package router

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/yuorei/video-server/app/adapter/infrastructure"
	"github.com/yuorei/video-server/app/adapter/presentation"
	"github.com/yuorei/video-server/app/application"
	flog "github.com/yuorei/video-server/app/driver/log"
	"github.com/yuorei/video-server/app/driver/newrelic"
	"github.com/yuorei/video-server/app/driver/sentry"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

const (
	defaultPort = "50051"
	httpAddr    = ":8081"
)

func NewRouter() {
	flog.NewLog()
	slog.Info("start server")

	sentry.SentryInit()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	// Setup metrics
	// srvMetrics := grpcprom.NewServerMetrics(
	// 	grpcprom.WithServerHandlingTimeHistogram(
	// 		grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
	// 	),
	// )
	// reg := prometheus.NewRegistry()
	// reg.MustRegister(srvMetrics)
	newrelicApp := newrelic.NewRelic()
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			nrgrpc.UnaryServerInterceptor(newrelicApp),
		),
		grpc.ChainStreamInterceptor(
			nrgrpc.StreamServerInterceptor(newrelicApp),
		),
	)

	infra := infrastructure.NewInfrastructure()
	app := application.NewApplication(infra)

	video_grpc.RegisterUserServiceServer(s, presentation.NewUserService(app))
	video_grpc.RegisterCommentServiceServer(s, presentation.NewCommentService(app))
	video_grpc.RegisterVideoServiceServer(s, presentation.NewVideoService(app))
	reflection.Register(s)

	// srvMetrics.InitializeMetrics(s)

	// Run gRPC server and HTTP server for Prometheus metrics
	var g run.Group
	g.Add(func() error {
		log.Printf("start gRPC server port: %v", port)
		return s.Serve(listener)
	}, func(err error) {
		s.GracefulStop()
		s.Stop()
	})

	// httpSrv := &http.Server{Addr: httpAddr}
	// g.Add(func() error {
	// 	m := http.NewServeMux()
	// 	m.Handle("/metrics", promhttp.HandlerFor(
	// 		reg,
	// 		promhttp.HandlerOpts{
	// 			EnableOpenMetrics: true,
	// 		},
	// 	))
	// 	httpSrv.Handler = m
	// 	log.Printf("start HTTP server for Prometheus metrics: %v", httpAddr)
	// 	return httpSrv.ListenAndServe()
	// }, func(error) {
	// 	if err := httpSrv.Close(); err != nil {
	// 		log.Printf("failed to stop HTTP server: %v", err)
	// 	}
	// })

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
