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
	errContainer := errors.NewErrContainer(fc)

	if fc.function.definition.Args != nil || fc.args != nil {
		err = fc.checkArgsCountMatch()
		if err != nil {
			// this is major error, we should break validation
			return nil, nil, errContainer.Wrap(err)
		}
		namedResult, namedExpressionListAst, err = fc.args.Check(env, validMngr)
		if err != nil {
			// this is major error, we should break validation
			return nil, nil, errContainer.Wrap(err)
		}
		for _, defArg := range fc.function.definition.Args {
			inputArg := namedResult.Get(defArg.VarName)
			if inputArg != defArg.VarType {
				errContainer.Add(errors.NewErrTypesMismatch(defArg, defArg.VarType, inputArg))
			}
			functionEnv.Set(defArg.VarName, inputArg)
		}
	}
	bodyAst, err := fc.function.body.Check(functionEnv, validMngr)
	if err != nil {
		return nil, nil, err
	}

	res := result.NewResult()
	for _, returnVar := range fc.function.definition.Returns {
		actualType, exists := functionEnv.Get(returnVar.VarName)
		if !exists {
			res.Add(returnVar.VarType)
		} else if actualType != returnVar.VarType {
			errContainer.Add(errors.NewErrTypesMismatch(returnVar, returnVar.VarType, actualType))
		} else {
			res.Add(actualType)
		}
	}
	if errContainer.NotEmpty() {
		return nil, nil, errContainer
	}

	functionAst := execAst.NewFunction(fc.function.id, fc.function.definition, bodyAst)
	functionCallAst := execAst.NewFunctionCall(fc.id, functionAst, namedExpressionListAst)
	return res, functionCallAst, nil
}

func (fc *FunctionCall) checkArgsCountMatch() error {
	definitionArgCount, inputArgCount := 0, 0
	if fc.function.definition.Args != nil {
		definitionArgCount = len(fc.function.definition.Args)
	}
	if fc.args != nil {
		inputArgCount = len(fc.args.exprs)
	}
	if definitionArgCount == inputArgCount {
		return nil
	}

	return errors.NewErrArgCountMismatch(fc, definitionArgCount, inputArgCount)
}
