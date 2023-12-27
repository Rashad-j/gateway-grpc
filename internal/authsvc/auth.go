package authsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/internal/config"
)

type AuthService interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type ServiceClient struct {
	// Client is the gRPC client for the auth service
	// TODO: Add the gRPC client for the auth service
	Client interface{}
}

func NewAuthServiceClient(cfg *config.Config) *ServiceClient {
	// TODO: Initialize the gRPC client for the auth service
	// TODO: Return the real gRPC client from the auth microservice
	return &ServiceClient{}
}
