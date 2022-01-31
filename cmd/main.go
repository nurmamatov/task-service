package main

import (
	"fmt"
	"khusniddin/task-servise/config"
	pb "khusniddin/task-servise/genproto"
	"khusniddin/task-servise/pkg/db"
	"khusniddin/task-servise/pkg/logger"
	"khusniddin/task-servise/service"
	"net"
	"time"

	// "github.com/gogo/protobuf/protoc-gen-gogo/grpc"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC3339))
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "task-servise")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
