package workflow

import (
	"context"
	"sync"
)

type WorkflowHandler interface {
	ExecuteActivity(ctx context.Context, definition ActivityDefinition, params any) (*ExecuteFuncResult, error)
}

type Engine struct {
	mu sync.Mutex

	handlers map[WorkflowName]WorkflowHandler
}

func NewEngine() *Engine {
	engine := &Engine{}

	go engine.run()

	return engine
}

func (e *Engine) run() {
}
