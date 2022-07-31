package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewStruct(name string, fields *Assignment) *Struct {
	return &Struct{
		key:    KeyStruct,
		name:   name,
		fields: fields,
	}
}

type Struct struct {
	id     int64
	key    NodeKey
	name   string
	fields *Assignment
}

func (s *Struct) ID() int64        { return s.id }
func (s *Struct) NodeKey() NodeKey { return s.key }

func (s *Struct) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	definition, _ := execMngr.MainPackage().StructDefinition(s.name)
	fields := make(map[string]object.Object)
	newResult := object.NewResult()
	execMngr.AddNextExec(s.fields, func() error {
		return s.fields.Exec(env, newResult, execMngr)
	})

	for i := range definition.Fields {
		ii := i
		execMngr.AddNextExec(s.fields.left[ii], func() error {
			fields[s.fields.left[ii].value] = newResult.ObjectList[ii]
			return nil
		})
	}

	execMngr.AddNextExec(s, func() error {
		result.Add(&object.ObjStruct{
			Fields: fields,
		})
		return nil
	})
	return nil
}
