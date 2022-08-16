package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
)

func NewPackage(id int64, name string) *Package {
	return &Package{
		id:        id,
		Name:      name,
		functions: make(map[string]*Function),
	}
}

type Package struct {
	Name      string
	id        int64
	functions map[string]*Function
}

func (p *Package) ID() int64            { return p.id }
func (p *Package) NodeKey() ast.NodeKey { return ast.KeyPackage }

func (p *Package) RegisterFunction(f *Function) error {
	if _, exists := p.functions[f.definition.Name]; exists {
		return errors.NewErrFunctionAlreadyExists(p, f.definition.Name)
	}
	p.functions[f.definition.Name] = f
	return nil
}

func (p *Package) Function(name string) (*Function, bool) {
	f, ok := p.functions[name]
	return f, ok
}
