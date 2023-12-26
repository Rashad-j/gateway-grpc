package authsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/internal/config"
)

func RegisterRoutes(r *gin.RouterGroup, cfg *config.Config) (*ServiceClient, error) {
	svc := &ServiceClient{
		Client: InitClient(cfg),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc, nil
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	// TODO: Implement the Register handler
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	// TODO: Implement the Login handler
}
