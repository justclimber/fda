package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewFunctionCall(id int64, function *Function, args *NamedExpressionList) *FunctionCall {
	return &FunctionCall{
		id:       id,
		function: function,
		args:     args,
	}
}

type FunctionCall struct {
	id       int64
	function *Function
	args     *NamedExpressionList
}

func (fc *FunctionCall) ID() int64            { return fc.id }
func (fc *FunctionCall) NodeKey() ast.NodeKey { return ast.KeyFunctionCall }

func (fc *FunctionCall) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	functionEnv := environment.NewEnclosedEnvironment(env)
	if fc.function.definition.Args != nil {
		namedResult := object.NewNamedResult()
		execMngr.AddNextExec(fc.args, func() error {
			return fc.args.Exec(functionEnv, namedResult, execMngr)
		})
		execMngr.AddNextExec(fc.args, func() error {
			for _, arg := range fc.function.definition.Args {
				inputArg := namedResult.Get(arg.VarName)
				functionEnv.Set(arg.VarName, inputArg)
			}
			return nil
		})
	}
	execMngr.AddNextExec(fc.function.body, func() error {
		return fc.function.body.Exec(functionEnv, execMngr)
	})
	execMngr.AddNextExec(fc, func() error {
		for _, returnVar := range fc.function.definition.Returns {
			returnVarObj, ok := functionEnv.Get(returnVar.VarName)
			if !ok {
				returnVarObj = getEmptyObjectByType(returnVar.VarType)
			}
			result.Add(returnVarObj)
		}
		return nil
	})
	return nil
}

// todo move to object helpers?
func getEmptyObjectByType(varType object.Type) object.Object {
	switch varType {
	case object.TypeInt:
		return &object.ObjInteger{
			Emptier: object.Emptier{Empty: true},
			Value:   0,
		}
	}
	return nil
}
