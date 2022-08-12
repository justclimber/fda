package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
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
		err = fc.checkArgsCountMatch()
		if err != nil {
			return nil, nil, err
		}
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
	validationErrorSet := errors.NewValidationErrorSet()
	for _, returnVar := range fc.function.definition.Returns {
		exists, matched := functionEnv.Check(returnVar.VarName, returnVar.VarType)
		if !exists {
			res.Add(returnVar.VarType)
		} else if !matched {
			validationErrorSet.Add(errors.NewValidationError(fc, errors.ErrorTypeMismatch))
		}
	}
	if !validationErrorSet.Empty() {
		return nil, nil, validationErrorSet
	}

	functionAst := execAst.NewFunction(fc.function.id, fc.function.definition, bodyAst)
	return res, execAst.NewFunctionCall(fc.id, functionAst, namedExpressionListAst), nil
}

func (fc *FunctionCall) checkArgsCountMatch() error {
	definitionArgCount, inputArgCount := 0, 0
	if fc.function.definition.Args != nil {
		definitionArgCount = len(fc.function.definition.Args)
	}
	if fc.args != nil {
		inputArgCount = len(fc.args.exprs)
	}

	if definitionArgCount != inputArgCount {
		return errors.NewRuntimeError(fc, errors.ErrorTypeArgumentsCountMismatch)
	}
	return nil
}
