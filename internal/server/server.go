package server

import (
	"context"

	"distri/internal/api/config"
)

type Server struct {
	cfg config.Config
}

func NewEngine(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	return nil
}
