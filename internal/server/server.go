package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rashad-j/grpc-gateway/internal/config"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Engine     *gin.Engine
	BaseRouter *gin.RouterGroup
	Version    string
}

func NewServer(version string) *Server {
	e := gin.New()
	r := e.Group(fmt.Sprintf("/%s", version))
	return &Server{
		Engine:     e,
		BaseRouter: r,
		Version:    version,
	}
}

func (s *Server) Serve(cfg *config.Config) error {
	listener, err := net.Listen("tcp", cfg.ServerAddr)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	srv := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: s.Engine,
	}

	log.Info().Str("addr", cfg.ServerAddr).Msg("starting gateway server")

	if cfg.TLSEnabled {
		if err := srv.ServeTLS(listener, cfg.TLSCertFile, cfg.TLSKeyFile); err != nil {
			return errors.Wrap(err, "failed to serve tls")
		}
	} else {
		if err := srv.Serve(listener); err != nil {
			return errors.Wrap(err, "failed to serve")
		}
	}
	return nil
}
