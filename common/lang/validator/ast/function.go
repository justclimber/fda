package ast

import (
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
	id         int64
	definition *object.FunctionDefinition
	body       *StatementsBlock
}

func (f *Function) ID() int64            { return f.id }
func (f *Function) NodeKey() ast.NodeKey { return ast.KeyFunction }

func (f *Function) Check(env ValidatorEnv, validMngr validationManager) (execAst.Stmt, error) {
	return nil, nil
}
