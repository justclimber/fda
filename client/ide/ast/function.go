package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
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

func (f *Function) Draw(r Renderer) {
	endFunc := r.StartSiblingNode(f)
	r.DrawFuncHeader(f.definition)
	f.body.Draw(r)
	r.DrawFuncBottom()
	endFunc()
}
