package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewPackage(id int64) *Package {
	return &Package{
		id:                  id,
		functionDefinitions: make(map[string]*object.FunctionDefinition),
	}
}

type Package struct {
	id                  int64
	functionDefinitions map[string]*object.FunctionDefinition
}

func (p *Package) ID() int64            { return p.id }
func (p *Package) NodeKey() ast.NodeKey { return ast.KeyPackage }

func (p *Package) RegisterFunctionDefinition(f *object.FunctionDefinition) {
	p.functionDefinitions[f.Name] = f
}

func (p *Package) FunctionDefinition(name string) (*object.FunctionDefinition, bool) {
	f, ok := p.functionDefinitions[name]
	return f, ok
}
