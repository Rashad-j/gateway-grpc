package server

import (
	"strconv"
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
