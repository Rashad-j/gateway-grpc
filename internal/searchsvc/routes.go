package searchsvc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/rpc/search"
)

func RegisterRoutes(r *gin.RouterGroup, svc SearchService) {
	routes := r.Group("/search")
	routes.Use(instrument)

	routes.GET("/:number", svc.search)
	routes.POST("/", svc.insert)
	routes.DELETE("/:number", svc.delete)
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
