package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewStructFieldAssignment(left []*StructFieldIdentifier, value Expr) *StructFieldAssignment {
	return &StructFieldAssignment{
		key:   KeyStructFieldAssignment,
		left:  left,
		value: value,
	}
}

type StructFieldAssignment struct {
	id    int64
	key   NodeKey
	left  []*StructFieldIdentifier
	value Expr
}

func (sf *StructFieldAssignment) NodeKey() NodeKey { return sf.key }
func (sf *StructFieldAssignment) ID() int64        { return sf.id }

func (sf *StructFieldAssignment) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(sf.value, func() error {
		return sf.value.Exec(env, result, execMngr)
	})
	res := object.NewResult()
	for i := range sf.left {
		ii := i
		execMngr.AddNextExec(sf.left[ii], func() error {
			return sf.left[ii].structIdentifier.Exec(env, res, execMngr)
		})
		execMngr.AddNextExec(sf.left[ii], func() error {
			objStruct := res.ObjectList[0].(*object.ObjStruct)
			objStruct.Fields[sf.left[ii].fieldName] = result.ObjectList[ii]
			result.Add(objStruct.Fields[sf.left[ii].fieldName])
			return nil
		})
	}
	return nil
}
