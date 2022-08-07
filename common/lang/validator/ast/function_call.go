package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
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

func (fc *FunctionCall) Exec(env *environment.Environment, validMngr validationManager) (*object.Result, execAst.Expr, error) {
	result := object.NewResult()
	functionEnv := environment.NewEnclosedEnvironment(env)
	var namedExpressionListAst *execAst.NamedExpressionList
	var err error
	var namedResult *object.NamedResult

	if fc.function.definition.Args != nil || fc.args != nil {
		// todo check count actual args and count args in definition
		namedResult, namedExpressionListAst, err = fc.args.Exec(env, validMngr)
		if err != nil {
			return nil, nil, err
		}
		validationErrorSet := errors.NewValidationErrorSet()
		for _, arg := range fc.function.definition.Args {
			inputArg := namedResult.Get(arg.VarName)
			if inputArg.Type() != arg.VarType {
				validationErrorSet.Add(errors.NewValidationError(arg, errors.ErrorTypeMismatch))
			}
			functionEnv.Set(arg.VarName, inputArg)
		}
		if !validationErrorSet.Empty() {
			return nil, nil, validationErrorSet
		}
	}
	bodyAst, err := fc.function.body.Exec(functionEnv, validMngr)
	if err != nil {
		return nil, nil, err
	}

	for _, returnVar := range fc.function.definition.Returns {
		// todo return vars check
		returnVarObj, ok := functionEnv.Get(returnVar.VarName)
		if !ok {
			returnVarObj = getEmptyObjectByType(returnVar.VarType)
		}
		result.Add(returnVarObj)
	}
	functionAst := execAst.NewFunction(fc.function.id, fc.function.definition, bodyAst)
	return result, execAst.NewFunctionCall(fc.id, functionAst, namedExpressionListAst), nil
}

// todo move to object helpers?
func getEmptyObjectByType(varType object.ObjectType) object.Object {
	switch varType {
	case object.TypeInt:
		return &object.ObjInteger{
			Emptier: object.Emptier{Empty: true},
			Value:   0,
		}
	}
	return nil
}
