package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewStruct(name string, fields *NamedExpressionList) *Struct {
	return &Struct{
		name:   name,
		fields: fields,
	}
}

type Struct struct {
	id     int64
	name   string
	fields *NamedExpressionList
}

func (s *Struct) ID() int64            { return s.id }
func (s *Struct) NodeKey() ast.NodeKey { return ast.KeyStruct }

func (s *Struct) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	definition, _ := execMngr.MainPackage().StructDefinition(s.name)
	fields := make(map[string]object.Object)
	newResult := object.NewNamedResult()
	execMngr.AddNextExec(s.fields, func() error {
		return s.fields.Exec(env, newResult, execMngr)
	})

	for name, _ := range definition.Fields {
		tName := name
		execMngr.AddNextExec(s.fields.exprs[tName], func() error {
			fields[tName] = newResult.ObjectList[tName]
			return nil
		})
	}

	execMngr.AddNextExec(s, func() error {
		result.Add(&object.ObjStruct{
			Name:   s.name,
			Fields: fields,
		})
		return nil
	})
	return nil
}
