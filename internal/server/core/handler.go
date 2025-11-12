package core

import (
	"context"
	"fmt"
)

type HandleFunc func(ctx context.Context, d *Distri) (res any, err error)

type Exec struct {
	EventVal any
}

func (e *Exec) Response(res any) {

}

type Handler struct {
	handleFunc HandleFunc

	execChan <-chan *Exec
}

func NewHandler(handleFunc HandleFunc) *Handler {
	return &Handler{
		handleFunc: handleFunc,
	}
}

func (h *Handler) Start(_ context.Context) error {
	for exec := range h.execChan {
		res, err := h.HandleExec(exec)
		if err != nil {
			exec.Response(err)

			continue
		}

		exec.Response(res)
	}

	return nil
}

// MapExec executes the function using reflection, so we can get all the information
// about possible signals, maybe crons, timeouts etc.
func (h *Handler) MapExec(ctx context.Context) error {
	distri := &Distri{
		runtime: runtimeMap,
	}

	h.handleFunc(ctx, distri)

	return nil
}

func (h *Handler) HandleExec(exec *Exec) (any, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	distri := &Distri{
		runtime: runtimeExec,
		event:   exec.EventVal,
	}

	res, err := h.handleFunc(ctx, distri)
	if err != nil {
		return nil, fmt.Errorf("handle with handle func: %w", err)
	}

	return res, err
}
