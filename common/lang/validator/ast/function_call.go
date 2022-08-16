package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/validator/result"
)

func NewFunctionCall(id int64, functionName string, packageName string, args *NamedExpressionList) *FunctionCall {
	return &FunctionCall{
		id:           id,
		functionName: functionName,
		packageName:  packageName,
		args:         args,
	}
}

type FunctionCall struct {
	id           int64
	functionName string
	packageName  string
	args         *NamedExpressionList
}

func (fc *FunctionCall) ID() int64            { return fc.id }
func (fc *FunctionCall) NodeKey() ast.NodeKey { return ast.KeyFunctionCall }

func (fc *FunctionCall) Check(env ValidatorEnv, validMngr validationManager) (*result.Result, execAst.Expr, error) {
	var namedExpressionListAst *execAst.NamedExpressionList
	var namedResult *result.NamedResult

	function, err := fc.getFunctionFromPackage(validMngr)
	if err != nil {
		return nil, nil, err
	}
	errContainer := errors.NewErrContainer(fc)

	if function.definition.Args != nil || fc.args != nil {
		if err = fc.checkArgsCountMatch(function); err != nil {
			// this is major error, we should break validation
			return nil, nil, errContainer.Wrap(err)
		}
		namedResult, namedExpressionListAst, err = fc.args.Check(env, validMngr)
		if err != nil {
			// this is major error, we should break validation
			return nil, nil, errContainer.Wrap(err)
		}
		for _, defArg := range function.definition.Args {
			inputArg := namedResult.Get(defArg.VarName)
			if inputArg != defArg.VarType {
				errContainer.Add(errors.NewErrTypesMismatch(defArg, defArg.VarType, inputArg))
			}
		}
	}
	bodyAst, hasError, err := function.GetCompiled()
	if err != nil {
		return nil, nil, err
	}
	if hasError {
		errContainer.Add(errors.NewErrCalledFunctionContainsErrors(fc, fc.functionName, fc.packageName))
	}

	res := result.NewResult()
	for _, returnVar := range function.definition.Returns {
		res.Add(returnVar.VarType)
	}
	if errContainer.NotEmpty() {
		return nil, nil, errContainer
	}

	functionAst := execAst.NewFunction(function.id, function.definition, bodyAst)
	functionCallAst := execAst.NewFunctionCall(fc.id, functionAst, namedExpressionListAst)
	return res, functionCallAst, nil
}

func (fc *FunctionCall) getFunctionFromPackage(validMngr validationManager) (*Function, error) {
	pkg, ok := validMngr.PackageByName(fc.packageName)
	if !ok {
		return nil, errors.NewErrPackageNotFound(fc, fc.packageName)
	}
	function, ok := pkg.Function(fc.functionName)
	if !ok {
		return nil, errors.NewErrFunctionNotFound(fc, fc.functionName, fc.packageName)
	}
	return function, nil
}

func (fc *FunctionCall) checkArgsCountMatch(function *Function) error {
	definitionArgCount, inputArgCount := 0, 0
	if function.definition.Args != nil {
		definitionArgCount = len(function.definition.Args)
	}
	if fc.args != nil {
		inputArgCount = len(fc.args.exprs)
	}
	if definitionArgCount == inputArgCount {
		return nil
	}

	return errors.NewErrArgCountMismatch(fc, definitionArgCount, inputArgCount)
}
