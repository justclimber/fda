package validator

import (
	"github.com/justclimber/fda/common/lang/executor/object"
	"github.com/justclimber/fda/common/lang/validator/ast"
)

type Environment struct {
	store map[string]object.Type
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{store: map[string]object.Type{}}
}

func (e *Environment) Set(name string, objType object.Type) {
	e.store[name] = objType
}

func (e *Environment) Check(name string, objType object.Type) bool {
	t, exists := e.Get(name)
	if !exists {
		return false
	}
	return t == objType
}

func (e *Environment) Get(name string) (object.Type, bool) {
	t, exists := e.store[name]
	if !exists && e.outer != nil {
		t, exists = e.outer.Get(name)
	}
	return t, exists
}

func (e *Environment) NewEnclosedEnvironment() ast.ValidatorEnv {
	env := NewEnvironment()
	env.outer = e
	return env
}
