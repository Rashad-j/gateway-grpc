package parsersvc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/rpc/parser"
	"github.com/stretchr/testify/assert"
)

type mockParserService struct {
}

func (m *mockParserService) parseJsonFilesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		&parser.JsonResponse{
			PersonList: []*parser.JsonResponse_Person{
				{
					FirstName:   "John",
					LastName:    "Doe",
					Birthday:    "01/01/2000",
					Address:     "123 Main St",
					PhoneNumber: "123-456-7890",
				},
			},
		},
	)
}

func (m *mockParserService) parseJsonFilesHandlerError(ctx *gin.Context) {
	// TODO
}

func (m *mockParserService) parseJsonFilesHandlerBadRequest(ctx *gin.Context) {
	// TODO
}

func parseJsonFilesHandlerTestSetup() (*gin.Engine, *gin.RouterGroup) {
	r := gin.Default()
	r.Use(gin.Recovery())

	routes := r.Group("/v1")
	RegisterRoutes(routes, &mockParserService{})

	return r, routes
}

func TestParseJsonFilesHandler(t *testing.T) {
	// Arrange
	r, _ := parseJsonFilesHandlerTestSetup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/parse/", nil)
	var response parser.JsonResponse

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "John", response.PersonList[0].FirstName)
	assert.Equal(t, "Doe", response.PersonList[0].LastName)
	assert.Equal(t, "01/01/2000", response.PersonList[0].Birthday)
	assert.Equal(t, "123 Main St", response.PersonList[0].Address)
	assert.Equal(t, "123-456-7890", response.PersonList[0].PhoneNumber)

}
