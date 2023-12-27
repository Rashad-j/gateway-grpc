package searchsvc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/rpc/search"
	"github.com/stretchr/testify/assert"
)

type mockSearchServiceClient struct {
	position int32
}

func (m *mockSearchServiceClient) search(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &search.SearchResponse{
		Position: m.position,
	})
}

func (m *mockSearchServiceClient) insert(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &search.InsertResponse{
		Position: m.position,
	})
}

func (m *mockSearchServiceClient) delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &search.DeleteResponse{
		Position: m.position,
	})
}

func TestServiceClientEndpoints(t *testing.T) {
	testCases := []struct {
		name     string
		method   string
		url      string
		expected interface{}
	}{
		{
			name:   "Search",
			method: "GET",
			url:    "/v1/search/123",
			expected: &search.SearchResponse{
				Position: 1,
			},
		},
		{
			name:   "Insert",
			method: "POST",
			url:    "/v1/search/",
			expected: &search.InsertResponse{
				Position: 1,
			},
		},
		{
			name:   "Delete",
			method: "DELETE",
			url:    "/v1/search/123",
			expected: &search.DeleteResponse{
				Position: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			gin.SetMode(gin.TestMode)
			router := gin.New()
			recorder := httptest.NewRecorder()

			client := &mockSearchServiceClient{
				position: 1,
			}
			RegisterRoutes(router.Group("/v1"), client)

			// Act
			req, _ := http.NewRequest(tc.method, tc.url, nil)
			router.ServeHTTP(recorder, req)

			// Assert
			assert.Equal(t, http.StatusOK, recorder.Code)
			err := json.Unmarshal(recorder.Body.Bytes(), tc.expected)
			assert.NoError(t, err)
		})
	}
}
