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

func (f *Function) Draw(r Renderer, slug string) {
	endFunc := r.StartSiblingNode(f, slug)
	r.StartContainerNode()
	r.DrawFuncHeader(f.definition)
	f.body.Draw(r, "function body")
	r.DrawFuncBottom()
	r.EndContainerNode()
	endFunc()
}
