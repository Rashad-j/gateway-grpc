package server

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func Instrument(ctx *gin.Context) {
	rid := xid.New().String()
	ctx.Writer.Header().Set("X-Request-Id", rid)
	// before request
	start := time.Now()
	log.Info().
		Str("request_id", rid).
		Str("userId", ctx.GetString("userId")).
		Str("method", ctx.Request.Method).
		Str("path", ctx.Request.URL.Path).
		Msg("request received")
	ctx.Next()
	// after request
	duration := time.Since(start)
	log.Info().
		Str("request_id", rid).
		Str("userId", ctx.GetString("userId")).
		Str("method", ctx.Request.Method).
		Str("path", ctx.Request.URL.Path).
		Str("duration", strconv.FormatInt(duration.Microseconds(), 10)).
		Msg("request completed")
}

func Authenticate(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		log.Error().Msg("authorization header is missing")
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
