package searchsvc

import (
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/rpc/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	searchClient search.SearchServiceClient
}

func NewServiceClient(client search.SearchServiceClient) *ServiceClient {
	return &ServiceClient{
		searchClient: client,
	}
}

func MakeServiceClient(cfg *config.Config) (search.SearchServiceClient, error) {
	conn, err := grpc.Dial(cfg.SearchAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}
	searchClient := search.NewSearchServiceClient(conn)

	return searchClient, nil
}
