package parsersvc

import (
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/rpc/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	parserClient parser.JsonParsingServiceClient
}

func NewServiceClient(parserClient parser.JsonParsingServiceClient) *ServiceClient {
	return &ServiceClient{
		parserClient: parserClient,
	}
}

func MakeServiceClient(cfg *config.Config) (parser.JsonParsingServiceClient, error) {
	conn, err := grpc.Dial(cfg.ParserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}
	client := parser.NewJsonParsingServiceClient(conn)

	return client, nil
}
