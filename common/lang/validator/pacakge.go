package validator

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewPackage() *Package {
	return &Package{
		structDefinitions:   make(map[string]*object.StructDefinition),
		functionDefinitions: make(map[string]*object.FunctionDefinition),
	}
}

type Package struct {
	id                  int64
	structDefinitions   map[string]*object.StructDefinition
	functionDefinitions map[string]*object.FunctionDefinition
}

func (p *Package) ID() int64            { return p.id }
func (p *Package) NodeKey() ast.NodeKey { return ast.KeyPackage }

func (p *Package) RegisterStructDefinition(s *object.StructDefinition) {
	p.structDefinitions[s.Name] = s
}

func (p *Package) StructDefinition(name string) (*object.StructDefinition, bool) {
	s, ok := p.structDefinitions[name]
	return s, ok
}

func (p *Package) RegisterFunctionDefinition(f *object.FunctionDefinition) {
	p.functionDefinitions[f.Name] = f
}

func (p *Package) FunctionDefinition(name string) (*object.FunctionDefinition, bool) {
	f, ok := p.functionDefinitions[name]
	return f, ok
}
