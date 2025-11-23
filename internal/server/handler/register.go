package handler

import (
	"distri/internal/server/communicate"

	"github.com/nats-io/nats.go/jetstream"
)

type (
	ConsumeRegisterHandler func(handler RemoteHandler)
	CreateRemoteHandler    func(param CreateRemoteHandlerParams) RemoteHandler

	CreateRemoteHandlerParams struct {
	}
)

func MakeRegisterRemoteHandlerConsumer(
	consumeRegisterHandler ConsumeRegisterHandler,
	createRemoteHandler CreateRemoteHandler,
) communicate.MessageHandler {
	return func(msg jetstream.Msg) {
		han := createRemoteHandler(CreateRemoteHandlerParams{})

		consumeRegisterHandler(han)
	}
}
