package core

import (
	"context"
	"errors"
	"fmt"
)

type HandleFunc func(ctx context.Context, d *Distri)

type Exec struct {
	EventVal any
}

type Handler struct {
	key        string
	handleFunc HandleFunc

	execChan    <-chan *Exec
	distriSlice []*Distri
}

func NewHandler(key string, handleFunc HandleFunc) *Handler {
	return &Handler{
		key:        key,
		handleFunc: handleFunc,
	}
}

func (h *Handler) Start(_ context.Context) error {
	for exec := range h.execChan {
		err := h.HandleExec(exec)
		if err != nil {
			continue
		}
	}

	return nil
}

// MapExec executes the function using reflection, so we can get all the information
// about possible signals, maybe crons, timeouts etc.
func (h *Handler) MapExec(ctx context.Context) {
	distri := &Distri{
		runtime: runtimeMap,
	}

	h.handleFunc(ctx, distri)

	h.distriSlice = append(h.distriSlice, distri)
}

func (h *Handler) HandleExec(exec *Exec) error {
	ctx, cancel := context.WithCancelCause(context.Background())

	distri := &Distri{
		runtime: runtimeExec,
		event:   exec.EventVal,
		ctx:     ctx,
		cancel:  cancel,
	}

	go func() {
		h.handleFunc(ctx, distri)

		// cancel with nil as this is successfully executed func
		cancel(nil)
	}()

	select {
	case <-ctx.Done():
		err := context.Cause(ctx)

		if err != nil && errors.Is(err, ErrDistriError) {
			return fmt.Errorf("context canceled with cause: %w", err)
		}

		return nil
	}
}
