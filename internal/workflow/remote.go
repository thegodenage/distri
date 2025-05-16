package workflow

import (
	"context"
	"fmt"
)

type RemoteActivityHandler interface {
	ExecuteActivity(ctx context.Context, params any) (*ExecuteFuncResult, error)
}

type RemoteHandler struct {
	// todo: for now is not parallel ready, sync.Map
	activitiesMap map[ActivityDefinition]RemoteActivityHandler
}

func (r *RemoteHandler) ExecuteActivity(ctx context.Context, definition ActivityDefinition, params any) (*ExecuteFuncResult, error) {
	v, ok := r.activitiesMap[definition]
	if !ok {
		return nil, fmt.Errorf("unknown activity definition: %s", definition.Name)
	}

	result, err := v.ExecuteActivity(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("execute activity %s: %w", definition.Name, err)
	}

	return result, nil
}
