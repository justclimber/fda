package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
)

func NewPackage(id int64, name string) *Package {
	return &Package{
		id:               id,
		Name:             name,
		functionsMap:     make(map[string]*Function),
		functionsOrdered: make([]*Function, 0),
	}
}

type Package struct {
	Name             string
	id               int64
	functionsMap     map[string]*Function
	functionsOrdered []*Function
}

func (p *Package) ID() int64            { return p.id }
func (p *Package) NodeKey() ast.NodeKey { return ast.KeyPackage }

func (p *Package) RegisterFunction(f *Function) error {
	if _, exists := p.functionsMap[f.definition.Name]; exists {
		return errors.NewErrFunctionAlreadyExists(p, f.definition.Name)
	}
	p.functionsMap[f.definition.Name] = f
	p.functionsOrdered = append(p.functionsOrdered, f)
	return nil
}

func (p *Package) Draw(r Renderer) {
	for _, function := range p.functionsOrdered {
		function.Draw(r)
	}
}
