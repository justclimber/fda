package ast

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewFunction(id int64, definition *object.FunctionDefinition, body *StatementsBlock) *Function {
	return &Function{
		id:         id,
		definition: definition,
		body:       body,
	}
}

type Function struct {
	id             int64
	definition     *object.FunctionDefinition
	body           *StatementsBlock
	compiled       bool
	compiledAst    *execAst.StatementsBlock
	compiledHasErr bool
}

func (f *Function) ID() int64            { return f.id }
func (f *Function) NodeKey() ast.NodeKey { return ast.KeyFunction }

func (f *Function) Compile(env ValidatorEnv, validMngr validationManager) error {
	var err error
	f.compiled = true
	for _, defArg := range f.definition.Args {
		env.Set(defArg.VarName, defArg.VarType)
	}
	for _, returnVar := range f.definition.Returns {
		env.Set(returnVar.VarName, returnVar.VarType)
	}
	f.compiledAst, err = f.body.Check(env, validMngr)
	if err != nil {
		f.compiledHasErr = true
		return err
	}
	return nil
}

func (f *Function) GetCompiled() (*execAst.StatementsBlock, bool, error) {
	if !f.compiled {
		return nil, false, fmt.Errorf("function %s#%s hadn't compiled", f.definition.PackageName, f.definition.Name)
	}
	return f.compiledAst, f.compiledHasErr, nil
}
