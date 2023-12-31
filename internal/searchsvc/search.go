package searchsvc

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/rpc/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SearchService interface {
	search(ctx *gin.Context)
	insert(ctx *gin.Context)
	delete(ctx *gin.Context)
}
type ServiceClient struct {
	searchClient search.SearchServiceClient
}

func NewServiceClient(client search.SearchServiceClient) *ServiceClient {
	return &ServiceClient{
		searchClient: client,
	}
}

func DialServiceClient(cfg *config.Config) (search.SearchServiceClient, error) {
	conn, err := grpc.Dial(cfg.SearchAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}
	searchClient := search.NewSearchServiceClient(conn)

	// test connection is successful
	_, err = searchClient.Search(context.Background(), &search.SearchRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}

	return searchClient, nil
}
