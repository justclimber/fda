package object

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewStructDefinition(name string, fields map[string]*VarAndType) *StructDefinition {
	return &StructDefinition{
		Name:   name,
		Fields: fields,
	}
}

type StructDefinition struct {
	id     int64
	Name   string
	Fields map[string]*VarAndType
}

func (sd *StructDefinition) ID() int64            { return sd.id }
func (sd *StructDefinition) NodeKey() ast.NodeKey { return ast.KeyStructDefinition }
