package main

import (
	"os"

	"github.com/rashad-j/grpc-gateway/internal/authsvc"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rashad-j/grpc-gateway/internal/parsersvc"
	"github.com/rashad-j/grpc-gateway/internal/searchsvc"
	"github.com/rashad-j/grpc-gateway/internal/server"
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

	// create server with version and base router
	s := server.NewServer(cfg.Version)
	s.BaseRouter.Use(server.Instrument)

	// gRPC client for the auth service
	authSvc, err := authsvc.RegisterRoutes(s.BaseRouter, cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to register auth routes")
		panic(err)
	}

	// gRPC client for the parser service
	if err := parsersvc.RegisterRoutes(s.BaseRouter, cfg, authSvc); err != nil {
		log.Error().Err(err).Msg("failed to register parser routes")
		panic(err)
	}

	// gRPC client for the search service
	if err := searchsvc.RegisterRoutes(s.BaseRouter, cfg, authSvc); err != nil {
		log.Error().Err(err).Msg("failed to register search routes")
		panic(err)
	}

	if err := s.Serve(cfg); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		panic(err)
	}
}
