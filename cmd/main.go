package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rashad-j/grpc-gateway/internal/authsvc"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/internal/parsersvc"
	"github.com/rashad-j/grpc-gateway/internal/searchsvc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// zerolog basic config
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})

	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	r := gin.New()
	v1 := r.Group("/v1")

	// gRPC client for the auth service
	authSvc, err := authsvc.RegisterRoutes(v1, cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to register auth routes")
		panic(err)
	}

	// gRPC client for the parser service
	if err := parsersvc.RegisterRoutes(v1, cfg, authSvc); err != nil {
		log.Error().Err(err).Msg("failed to register parser routes")
		panic(err)
	}

	// gRPC client for the search service
	if err := searchsvc.RegisterRoutes(v1, cfg, authSvc); err != nil {
		log.Error().Err(err).Msg("failed to register search routes")
		panic(err)
	}

	l, err := net.Listen("tcp", cfg.ServerAddr)
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		panic(err)
	}
	srv := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	log.Info().Str("addr", cfg.ServerAddr).Msg("starting gateway server")

	if cfg.TLSEnabled {
		if err := srv.ServeTLS(l, cfg.TLSCertFile, cfg.TLSKeyFile); err != nil {
			log.Error().Err(err).Msg("failed to serve")
			panic(err)
		}
	} else {
		if err := srv.Serve(l); err != nil {
			log.Error().Err(err).Msg("failed to serve")
			panic(err)
		}
	}
}
