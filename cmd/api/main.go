package main

import (
	"context"
	"os"

	"github.com/nats-io/nats.go"

	"distri/internal/environment"
	"distri/internal/server"
	"distri/internal/server/config"
)

var (
	configOptions = []config.Option{
		config.NatsAddress(os.Getenv(server.EnvAddress)),
	}
	devConfigOptions = []config.Option{
		config.NatsAddress(nats.DefaultURL),
	}
)

func main() {
	cfg := config.NewConfig(configOptions...)

	if environment.IsDevelopment(os.Getenv(environment.Env)) {
		cfg = config.NewConfig(devConfigOptions...)
	}

	if ok, err := cfg.IsValid(); !ok {
		panic(err.Error())
	}

	if err := server.NewEngine(cfg).Run(context.Background()); err != nil {
		panic(err.Error())
	}
}
