package main

import (
	"context"
	"distri/internal/choreography"
)

var da = &domainActivities{}

type domainActivities struct {
}

func (d *domainActivities) ExecuteAsyncActivity(ctx context.Context, param string) (any, error) {
	return nil, nil
}

func main() {
	conductor := choreography.NewConductor("AsyncBusinessProcess")
	conductor.Activity(da.ExecuteAsyncActivity)
}
