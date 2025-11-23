package server

import (
	"context"
	"distri/internal/server/communicate"
	"distri/internal/server/config"
	"distri/internal/server/handler"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Engine struct {
	cfg config.Config
}

func NewEngine(cfg config.Config) *Engine {
	return &Engine{
		cfg: cfg,
	}
}

// 3 components:
//	messaging,
//	task management,
//	task scheduling & response management

func (s *Engine) Run(ctx context.Context) error {
	conn, err := nats.Connect(s.cfg.NatsURL)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS: %w", err)
	}

	js, err := jetstream.New(conn)
	if err != nil {
		return fmt.Errorf("failed to create Jetstream: %w", err)
	}

	communicationHandler := communicate.Handler(conn, js)

	manager := handler.NewManager()

	registerHandlerConsumer := handler.MakeRegisterRemoteHandlerConsumer(
		manager.RegisterRemoteHandler,
		func(param handler.CreateRemoteHandlerParams) handler.RemoteHandler {
			return handler.NewRemoteHandler()
		},
	)

	consumerCfg := communicate.MakeConsumerCfg("register_handler", registerHandlerConsumer)

	err = communicationHandler.Consume(ctx, consumerCfg)
	if err != nil {
		return fmt.Errorf("failed to register handler: %w", err)
	}

	return nil
}
