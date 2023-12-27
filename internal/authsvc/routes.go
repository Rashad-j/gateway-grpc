package authsvc

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, svc AuthService) {
	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	// TODO: Implement the Register handler
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	// TODO: Implement the Login handler
}
