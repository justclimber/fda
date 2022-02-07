package ecs

type nestedDebugger interface {
	LogF(method string, str string, args ...interface{})
}

type emptyDebugger struct{}

func (e *emptyDebugger) LogF(_ string, _ string, _ ...interface{}) {}
