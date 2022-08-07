package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewStructFieldIdentifier(fieldName string, structIdentifier *Identifier) *StructFieldIdentifier {
	return &StructFieldIdentifier{
		fieldName:        fieldName,
		structIdentifier: structIdentifier,
	}
}

type StructFieldIdentifier struct {
	id               int64
	fieldName        string
	structIdentifier *Identifier
}

func (sf *StructFieldIdentifier) ID() int64            { return sf.id }
func (sf *StructFieldIdentifier) NodeKey() ast.NodeKey { return ast.KeyStructFieldIdentifier }
