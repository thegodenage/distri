package main

import (
	"context"
	"distri/internal/server/core"
	"fmt"
)

// main here it's only for concept of the app, not an actual usable command.
func main() {
	cfg := &core.AppConfiguration{}

	handler := core.NewHandler("test", func(ctx context.Context, d *core.Distri) {
		event := d.OnEvent("test")

		v, err := core.Map[string](event)
		d.Error(err)

		res, err := d.Func(func() (any, error) {
			return *v + ": processed", nil
		})
		d.Error(err)

		d.SendEvent("test1", res)
	})

	handler.MapExec(context.Background())

	exec := &core.Exec{
		EventVal: "to jest po prostu jakis testowy event",
	}

	err := handler.HandleExec(exec)
	fmt.Printf("err: %e", err)

	app := core.NewApp(cfg, handler)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
