package searchsvc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/authsvc"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/rpc/search"
)

func RegisterRoutes(r *gin.RouterGroup, cfg *config.Config, authSvc *authsvc.ServiceClient) error {
	// TODO: implement retry logic on connection failure
	parserClient, err := MakeServiceClient(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to dial server")
	}
	svc := NewServiceClient(parserClient)

	routes := r.Group("/search")
	routes.Use(authSvc.Authenticate)
	routes.Use(instrument)

	routes.GET("/:number", svc.search)
	routes.POST("/", svc.insert)
	routes.DELETE("/:number", svc.delete)

	return nil
}

func (sc *ServiceClient) search(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := sc.searchClient.Search(ctx, &search.SearchRequest{
		Number: int32(number),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, r)
}

func (sc *ServiceClient) insert(ctx *gin.Context) {
	var req search.InsertRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := sc.searchClient.Insert(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, r)
}

func (sc *ServiceClient) delete(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := sc.searchClient.Delete(ctx, &search.DeleteRequest{
		Number: int32(number),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, r)
}
