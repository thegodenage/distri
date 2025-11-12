package core

import (
	"context"
	"errors"
	"fmt"
)

type runtime string

const (
	// runtimeMap means that given distri func was executed during map phase
	runtimeMap = "runtime_map"
	// runtimeExec means that given distri func was executed during execution
	runtimeExec = "runtime_exec"
)

var (
	ErrDistriError = errors.New("distri registered error")
)

type LifecycleEvent struct {
	key string
}

// Distri is a bread and butter of distri. Use it only in the handle func to register functions, check for top level errors
// and use its lifecycle functions to send events or signals. Everything will be handled by core node and scheduled
// to be executed somewhere.
type Distri struct {
	runtime runtime
	ctx     context.Context
	cancel  context.CancelCauseFunc

	lifecycles []any

	event any
}

// OnEvent set up an event listener in given handler workflow
// that awaits event with given key, if event happens it executes the workflow
// and returns that event with this function.
func (d *Distri) OnEvent(key string) any {
	if d.runtime == runtimeMap {
		d.lifecycles = append(d.lifecycles, &LifecycleEvent{
			key: key,
		})

		return nil
	}

	return d.event
}

func (d *Distri) SendEvent(key string, val any) any {
	if d.runtime == runtimeMap {
		return nil
	}

	// send event here
	return nil
}

func (d *Distri) Func(f func() (any, error)) (any, error) {
	if d.runtime == runtimeMap {
		return nil, nil
	}

	select {
	case <-d.ctx.Done():
		return nil, d.ctx.Err()
	default:
		return f()
	}
}

// Error should be executed whenever we want to handle error just inside the handle func,
// because in map phase, we want just to ignore it atp. It will send the error higher and cancel the execution.
func (d *Distri) Error(err error) {
	if d.runtime == runtimeExec && err != nil {
		d.cancel(fmt.Errorf("%w: %s", ErrDistriError, err.Error()))

		return
	}
}
