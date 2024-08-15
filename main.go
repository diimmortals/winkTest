package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync/atomic"

	pb "wink/proto"

	"google.golang.org/grpc"
)

type Config struct {
	CDNHost string
}

func LoadConfig() *Config {
	return &Config{
		CDNHost: getEnv("CDN_HOST", "default-cdn-host"), // Значение по умолчанию
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

type server struct {
	pb.UnimplementedBalancerServer
	counter int32
	config  *Config
}

func (s *server) Redirect(ctx context.Context, req *pb.RedirectRequest) (*pb.RedirectResponse, error) {
	url := req.GetVideo()

	count := atomic.AddInt32(&s.counter, 1)

	if count%10 == 0 {
		return &pb.RedirectResponse{Url: url}, nil
	}

	cdnUrl := fmt.Sprintf("http://%s/%s", s.config.CDNHost, url[7:])

	return &pb.RedirectResponse{Url: cdnUrl}, nil

}

func main() {
	// Загружаем конфигурацию
	config := LoadConfig()

	// Создаем gRPC сервер
	lis, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис
	pb.RegisterBalancerServer(grpcServer, &server{config: config})

	log.Println("Starting gRPC server on port 443...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
