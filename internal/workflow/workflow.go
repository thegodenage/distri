package workflow

import "time"

type WorkflowName string

type WorkflowDefinition struct {
	WorkflowName          WorkflowName
	CreatedAt             time.Time
	ActivitiesDefinitions []ActivityDefinition
	Version               int
}

type ActivityDefinition struct {
	Name string
}
