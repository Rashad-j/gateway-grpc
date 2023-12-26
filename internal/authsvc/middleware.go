package authsvc

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *ServiceClient) Authenticate(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// TODO: Call the auth service to validate the token
	// TODO: If the token is invalid, abort the request with status code 401

	// set fake userId for now
	ctx.Set("userId", "rambo")

	ctx.Next()
}
