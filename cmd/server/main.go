package main

import (
	"context"

	"distri/internal/server/core"
)

// main here it's only for concept of the app, not an actual usable command.
func main() {
	cfg := &core.AppConfiguration{}

	app := core.NewApp(cfg)

	handler := core.NewHandler(func(ctx context.Context, d *core.Distri) (res any, err error) {
		event := d.OnEvent("test")

		return event, nil
	})

	handler.MapExec(context.Background())

	exec := &core.Exec{
		EventVal: "to jest po prostu jakis testowy event",
	}

	handler.HandleExec(exec)

	app.SetHandler(handler)
}
