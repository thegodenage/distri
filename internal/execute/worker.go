package execute

import (
	"context"
	"errors"
	"reflect"
)

var (
	ErrFirstParamNotCtx = errors.New("first param is not a execute.ctx")
)

type workflowMetadata struct {
	fn            reflect.Value
	fnParamsTypes []reflect.Type
	fnReturnTypes []reflect.Type
	actions       []any
}

type ContextBasedWorker struct {
	workflows []*workflowMetadata
}

func (c *ContextBasedWorker) RegisterWorkflow(workflows ...any) error {
	for _, workflow := range workflows {
		metadata := &workflowMetadata{}

		metadata.fn = reflect.ValueOf(workflow)

		fnType := metadata.fn.Type()

		if fnType.Kind() != reflect.Func {
			return errors.New("workflow must be a function")
		}

		for i := 0; i < fnType.NumIn(); i++ {
			metadata.fnParamsTypes = append(metadata.fnParamsTypes, fnType.In(i))
		}

		if len(metadata.fnParamsTypes) == 0 {
			return ErrFirstParamNotCtx
		}

		firstParamType := reflect.TypeOf(metadata.fnParamsTypes[0])
		if firstParamType != reflect.TypeOf(Context{}) {
			return ErrFirstParamNotCtx
		}

		for i := 0; i < fnType.NumOut(); i++ {
			metadata.fnReturnTypes = append(metadata.fnReturnTypes, fnType.Out(i))
		}

		c.workflows = append(c.workflows, metadata)
	}

	return nil
}

func (c *ContextBasedWorker) ExecuteAction(ctx context.Context, name string, data []byte) (any, error) {
	return nil, nil
}
