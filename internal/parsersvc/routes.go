package parsersvc

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/rpc/parser"
)

func RegisterRoutes(r *gin.RouterGroup, svc ParserService) error {
	routes := r.Group("/parse")
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
