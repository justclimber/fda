package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
	"github.com/justclimber/fda/common/lang/validator/result"
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

func (fc *FunctionCall) Check(env ValidatorEnv, validMngr validationManager) (*result.Result, execAst.Expr, error) {
	functionEnv := env.NewEnclosedEnvironment()
	var namedExpressionListAst *execAst.NamedExpressionList
	var err error
	var namedResult *result.NamedResult

	if fc.function.definition.Args != nil || fc.args != nil {
		// todo check count actual args and count args in definition
		namedResult, namedExpressionListAst, err = fc.args.Check(env, validMngr)
		if err != nil {
			return nil, nil, err
		}
		validationErrorSet := errors.NewValidationErrorSet()
		for _, arg := range fc.function.definition.Args {
			inputArg := namedResult.Get(arg.VarName)
			if inputArg != arg.VarType {
				validationErrorSet.Add(errors.NewValidationError(arg, errors.ErrorTypeMismatch))
			}
			functionEnv.Set(arg.VarName, inputArg)
		}
		if !validationErrorSet.Empty() {
			return nil, nil, validationErrorSet
		}
	}
	bodyAst, err := fc.function.body.Check(functionEnv, validMngr)
	if err != nil {
		return nil, nil, err
	}

	res := result.NewResult()
	for _, returnVar := range fc.function.definition.Returns {
		// todo return vars check
		returnVarObj, ok := functionEnv.Get(returnVar.VarName)
		if !ok {
			returnVarObj = returnVar.VarType
		}
		res.Add(returnVarObj)
	}
	functionAst := execAst.NewFunction(fc.function.id, fc.function.definition, bodyAst)
	return res, execAst.NewFunctionCall(fc.id, functionAst, namedExpressionListAst), nil
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
