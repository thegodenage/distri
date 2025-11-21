package core

import (
	"context"
	"errors"
)

type ActivityFunc func(ctx context.Context, maybe Maybe) Maybe

type (
	event             any
	phase             string
	workflowName      string
	phaseActivityFunc func(ctx context.Context, f ActivityFunc, maybe Maybe) Maybe
)

const (
	// phaseMap means that given distri func was executed during map phase
	phaseMap = "phase_map"
	// runtimeExec means that given distri func was executed during execution
	phaseExec = "phase_exec"
)

// Distri is a bread and butter of distri. Use it only in the handle func to register functions, check for top level errors
// and use its lifecycle functions to send events or signals. Everything will be handled by core node and scheduled
// to be executed somewhere.
type Distri struct {
	uniqueWorkflowName workflowName
	phaseActivityFunc  phaseActivityFunc
}

func NewDistri(name workflowName, activityFunc phaseActivityFunc) *Distri {
	return &Distri{
		uniqueWorkflowName: name,
		phaseActivityFunc:  activityFunc,
	}
}

func (d *Distri) NewActivity(ctx context.Context, f ActivityFunc, maybe Maybe) Maybe {
	if f == nil {
		_, cancel := context.WithCancelCause(ctx)

		err := errors.New("function passed to activityInfo is nil")

		cancel(err)

		return MaybeWithErr(err)
	}

	return d.phaseActivityFunc(ctx, f, maybe)
}

func (d *Distri) Done(ctx context.Context, param Maybe) {
	if param.Err != nil {

	}
}
