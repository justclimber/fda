package fdalang

import (
	"fmt"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	return &Environment{
		store:             make(map[string]Object),
		structDefinitions: make(map[string]*AstStructDefinition),
		enumDefinitions:   make(map[string]*AstEnumDefinition),
	}
}

type Environment struct {
	store             map[string]Object
	structDefinitions map[string]*AstStructDefinition
	enumDefinitions   map[string]*AstEnumDefinition
	outer             *Environment
}

func (e *Environment) Store() map[string]Object {
	return e.store
}
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) RegisterStructDefinition(s *AstStructDefinition) error {
	if _, exists := e.structDefinitions[s.Name]; exists {
		return fmt.Errorf("struct '%s' already defined in this scope", s.Name)
	}
	e.structDefinitions[s.Name] = s

	return nil
}

func (e *Environment) RegisterEnumDefinition(ed *AstEnumDefinition) error {
	if _, exists := e.enumDefinitions[ed.Name]; exists {
		return fmt.Errorf("enum '%s' already defined in this scope", ed.Name)
	}
	e.enumDefinitions[ed.Name] = ed

	return nil
}

func (e *Environment) StructDefinition(name string) (*AstStructDefinition, bool) {
	s, ok := e.structDefinitions[name]

	if !ok && e.outer != nil {
		s, ok = e.outer.StructDefinition(name)
	}

	return s, ok
}

func (e *Environment) EnumDefinition(name string) (*AstEnumDefinition, bool) {
	ed, ok := e.enumDefinitions[name]

	if !ok && e.outer != nil {
		ed, ok = e.outer.EnumDefinition(name)
	}

	return ed, ok
}
