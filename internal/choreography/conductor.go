package choreography

import (
	"fmt"
	"reflect"
	"sync"
)

type functionMetadata struct {
	fn       reflect.Value
	fnParams []reflect.Type
	fnReturn []reflect.Type
}

type Conductor struct {
	processName string
	functions   []functionMetadata

	mu sync.Mutex
}

func NewConductor(processName string) *Conductor {
	return &Conductor{
		processName: processName,
	}
}

func (c *Conductor) Activity(f any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	fnMetadata := functionMetadata{}

	fnMetadata.fn = reflect.ValueOf(f)

	fnType := fnMetadata.fn.Type()

	if fnType.Kind() != reflect.Func {
		return fmt.Errorf(fmt.Sprintf("expected a function, got %s", fnType.Kind()))
	}

	for i := 0; i < fnType.NumIn(); i++ {
		fnMetadata.fnParams = append(fnMetadata.fnReturn, fnType.In(i))
	}

	for i := 0; i < fnType.NumOut(); i++ {
		fnMetadata.fnReturn = append(fnMetadata.fnParams, fnType.In(i))
	}

	c.functions = append(c.functions, fnMetadata)

	return nil
}
