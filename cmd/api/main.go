package main

import (
	"context"
	"os"

	"distri/internal/api"
	"distri/internal/api/config"
	"distri/internal/environment"
	"distri/internal/server"
)

var (
	configOptions = []config.Option{
		config.Address(os.Getenv(api.EnvAddress)),
	}
	devConfigOptions = []config.Option{
		config.Address("8080"),
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
