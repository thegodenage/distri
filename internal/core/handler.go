package core

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Exec struct {
	WorkflowKey      string
	ActivityKey      string
	Param            any
	CachedResultsMap map[activityName]*CachedResult
	D                *Distri
}

type CachedResult struct {
	Maybe Maybe
}

type Handler struct {
	key       string
	workflows map[workflowName]*workflowInfo
}

type workflowInfo struct {
	fun        any
	t          reflect.Type
	activities []activityInfo
}

func NewHandler(key string) *Handler {
	return &Handler{
		key:       key,
		workflows: make(map[workflowName]*workflowInfo),
	}
}

func (h *Handler) RegisterWorkflow(key string, f any) error {
	t := reflect.TypeOf(f)

	if t.Kind() != reflect.Func {
		return errors.New("f is not a function")
	}

	if t.NumIn() == 2 || t.NumIn() == 3 {
		if !t.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
			return fmt.Errorf("first param must be context.Context")
		}

		if t.In(1) != reflect.TypeOf(&Distri{}) {
			return fmt.Errorf("second param must be *Distri")
		}

		h.workflows[workflowName(key)] = &workflowInfo{
			fun: f,
			t:   reflect.TypeOf(f),
		}

		return nil
	}

	return fmt.Errorf("invalid workflow function")
}

func (h *Handler) Run(ctx context.Context) error {
	for key, workflow := range h.workflows {
		fun := reflect.ValueOf(workflow.fun)

		distri := NewDistri(key, h.activityMapDecorator(key))

		if workflow.t.NumIn() == 2 {
			fun.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(distri),
			})
		}

		paramType := reflect.TypeOf(workflow.fun).In(2)

		if workflow.t.NumIn() == 3 {
			fun.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(distri),
				reflect.Zero(paramType),
			})
		}
	}

	return nil
}

func (h *Handler) activityMapDecorator(workflowName workflowName) phaseActivityFunc {
	return func(ctx context.Context, f ActivityFunc, maybe Maybe) Maybe {
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

		if strings.Contains(funcName, ".func") {
			return MaybeWithErr(errors.New("cannot use anonymous function as activityInfo"))
		}

		funcNameParts := strings.Split(funcName, ".")

		h.workflows[workflowName].activities = append(h.workflows[workflowName].activities, activityInfo{
			name: activityName(fmt.Sprintf("%s_%s", workflowName, funcNameParts[len(funcNameParts)-1])),
		})

		return Maybe{}
	}
}

type (
	activityName string

	activityInfo struct {
		name activityName
	}
)

func (h *Handler) Execute(ctx context.Context, e *Exec) error {
	workflow, ok := h.workflows[workflowName(e.WorkflowKey)]
	if !ok {
		return errors.New("workflow not found")
	}

	fun := reflect.ValueOf(workflow.fun)

	if workflow.t.NumIn() == 2 {
		fun.Call([]reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(e.D),
		})
	}

	if workflow.t.NumIn() == 3 {
		fun.Call([]reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(e.D),
			reflect.ValueOf(e.Param),
		})
	}

	return errors.New("invalid workflow function")
}
