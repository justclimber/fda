package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewStructFieldCall(fieldName string, structExpr Expr) *StructFieldCall {
	return &StructFieldCall{
		fieldName:  fieldName,
		structExpr: structExpr,
	}
}

type StructFieldCall struct {
	id         int64
	fieldName  string
	structExpr Expr
}

func (sf *StructFieldCall) ID() int64        { return sf.id }
func (sf *StructFieldCall) NodeKey() NodeKey { return KeyStructFieldCall }

func (sf *StructFieldCall) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	res := object.NewResult()
	execMngr.AddNextExec(sf.structExpr, func() error {
		return sf.structExpr.Exec(env, res, execMngr)
	})
	execMngr.AddNextExec(sf.structExpr, func() error {
		objStruct := res.ObjectList[0].(*object.ObjStruct)
		result.Add(objStruct.Fields[sf.fieldName])
		return nil
	})
	return nil
}
