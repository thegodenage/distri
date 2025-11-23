package communicate

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type (
	MessageHandler jetstream.MessageHandler

	InHandlerSubCfg struct {
		StreamName      string
		Subjects        []string
		ConsumerDurable string
		ConsumerName    string
		Consume         MessageHandler
	}

	InHandler struct {
		conn *nats.Conn
		js   jetstream.JetStream
	}
)

func Handler(
	conn *nats.Conn,
	js jetstream.JetStream,
) *InHandler {
	return &InHandler{
		conn: conn,
		js:   js,
	}
}

func (ih *InHandler) Consume(ctx context.Context, cfg InHandlerSubCfg) error {
	s, err := ih.js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     cfg.StreamName,
		Subjects: cfg.Subjects,
	})
	if err != nil {
		return fmt.Errorf("create stream: %w", err)
	}

	c, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   cfg.ConsumerDurable,
		Name:      cfg.ConsumerName,
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		return fmt.Errorf("create consumer: %w", err)
	}

	conns, err := c.Consume(func(msg jetstream.Msg) {
		cfg.Consume(msg)

		_ = msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("consume stream: %w", err)
	}

	defer conns.Stop()

	return nil
}
