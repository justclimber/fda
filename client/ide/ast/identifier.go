package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewIdentifier(id int64, name string) *Identifier {
	return &Identifier{
		id:   id,
		name: name,
	}
}

type Identifier struct {
	id   int64
	name string
}

func (i *Identifier) ID() int64            { return i.id }
func (i *Identifier) NodeKey() ast.NodeKey { return ast.KeyIdentifier }