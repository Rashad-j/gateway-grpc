package authsvc

import "github.com/rashad-j/grpc-gateway/internal/config"

type ServiceClient struct {
	// Client is the gRPC client for the auth service
	// TODO: Add the gRPC client for the auth service
	Client interface{}
}

func InitClient(cfg *config.Config) *ServiceClient {
	// TODO: Initialize the gRPC client for the auth service
	// TODO: Return the real gRPC client from the auth microservice
	return &ServiceClient{}
}
