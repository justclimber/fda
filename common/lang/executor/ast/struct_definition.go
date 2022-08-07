package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewStructDefinition(name string, fields map[string]*object.VarAndType) *StructDefinition {
	return &StructDefinition{
		Name:   name,
		Fields: fields,
	}
}

type StructDefinition struct {
	id     int64
	Name   string
	Fields map[string]*object.VarAndType
}

func (sd *StructDefinition) ID() int64            { return sd.id }
func (sd *StructDefinition) NodeKey() ast.NodeKey { return ast.KeyStructDefinition }
