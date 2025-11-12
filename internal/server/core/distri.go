package core

type runtime string

const (
	// runtimeMap means that given distri func was executed during map phase
	runtimeMap = "runtime_map"
	// runtimeExec means that given distri func was executed during execution
	runtimeExec = "runtime_exec"
)

type LifecycleEvent struct {
	key string
}

type Distri struct {
	runtime runtime

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
