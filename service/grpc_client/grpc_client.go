package grpcClient

import (
	"khusniddin/task-servise/config"
)

// GrpcClientI ...
type GrpcClientI interface {
}

// GrpcClient
type grpcClient struct {
	cfg        config.Config
	connection map[string]interface{}
}

//New ...
func New(cfg config.Config) (*grpcClient, error) {
	return &grpcClient{
		cfg:        cfg,
		connection: map[string]interface{}{},
	}, nil
}
