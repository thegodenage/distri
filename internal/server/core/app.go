package core

import (
	"errors"
	"fmt"
)

type AppConfiguration struct {
}

type App struct {
	cfg      *AppConfiguration
	handlers []*Handler
}

func NewApp(cfg *AppConfiguration, handlers ...*Handler) *App {
	return &App{
		cfg:      cfg,
		handlers: handlers,
	}
}

func (a *App) Start() error {
	if err := a.validateHandlers(); err != nil {
		return fmt.Errorf("validate handlers: %w", err)
	}

	return nil
}

func (a *App) validateHandlers() error {
	handlersMap := make(map[string]*Handler)

	for _, handler := range a.handlers {
		_, ok := handlersMap[handler.key]
		if ok {
			return errors.New("handler with given key already exists")
		}

		handlersMap[handler.key] = handler
	}

	return nil
}
