package parsersvc

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/rpc/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ParserService interface {
	parseJsonFilesHandler(*gin.Context)
}

type ServiceClient struct {
	parserClient parser.JsonParsingServiceClient
}

func NewServiceClient(parserClient parser.JsonParsingServiceClient) *ServiceClient {
	return &ServiceClient{
		parserClient: parserClient,
	}
}

func DialServiceClient(cfg *config.Config) (parser.JsonParsingServiceClient, error) {
	conn, err := grpc.Dial(cfg.ParserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}
	client := parser.NewJsonParsingServiceClient(conn)

	// test connection is successful
	_, err = client.ParseJsonFiles(context.Background(), &parser.EmptyRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}

	return client, nil
}
