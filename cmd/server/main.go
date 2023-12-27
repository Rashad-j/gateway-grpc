package main

import (
	"os"

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
	s.BaseRouter.Use(server.Authenticate)

	// gRPC client for the parser service
	parser, err := parsersvc.DialServiceClient(cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to dial parser service")
		panic(err)
	}
	parserClient := parsersvc.NewServiceClient(parser)
	parsersvc.RegisterRoutes(s.BaseRouter, parserClient)

	// gRPC client for the search service
	searcher, err := searchsvc.DialServiceClient(cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to dial search service")
		panic(err)
	}
	searchClient := searchsvc.NewServiceClient(searcher)
	searchsvc.RegisterRoutes(s.BaseRouter, searchClient)

	// start server
	if err := s.Serve(cfg); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		panic(err)
	}
}
