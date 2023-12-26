package parsersvc

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/authsvc"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/rpc/parser"
)

func RegisterRoutes(r *gin.RouterGroup, cfg *config.Config, authSvc *authsvc.ServiceClient) error {
	// TODO: implement retry logic on connection failure
	parserClient, err := MakeServiceClient(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to dial server")
	}
	svc := NewServiceClient(parserClient)

	routes := r.Group("/parse")
	routes.Use(authSvc.Authenticate)
	routes.Use(instrument)

	routes.GET("/", svc.parseJsonFilesHandler)

	return nil
}

func (sc *ServiceClient) parseJsonFilesHandler(c *gin.Context) {
	response, err := sc.parserClient.ParseJsonFiles(context.Background(), &parser.EmptyRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
