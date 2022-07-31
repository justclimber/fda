package environment

import (
	"encoding/json"
	"fmt"

	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]object.Object),
	}
}

type Environment struct {
	store map[string]object.Object
	outer *Environment
}

func (e *Environment) Store() map[string]object.Object {
	return e.store
}
func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}

func (e *Environment) Print() {
	fmt.Println("Env content:")
	for k, v := range e.store {
		fmt.Printf("%s: %s\n", k, v.Inspect())
	}
}

func (e *Environment) GetVarsAsJson() ([]byte, error) {
	varMap := make(map[string]string)
	for k, v := range e.store {
		varMap[k] = v.Inspect()
	}
	return json.Marshal(varMap)
}

func (e *Environment) ToStrings() []string {
	result := make([]string, 0)
	for k, v := range e.store {
		result = append(result, fmt.Sprintf("%s: %s\n", k, v.Inspect()))
	}
	return result
}

func (e *Environment) Keys() []string {
	keys := make([]string, len(e.store))

	i := 0
	for k := range e.store {
		keys[i] = k
		i++
	}
	return keys
}
