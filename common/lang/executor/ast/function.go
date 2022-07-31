package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewFunction(stmtsBlock *StatementsBlock) *Function {
	return &Function{
		key:             KeyFunction,
		statementsBlock: stmtsBlock,
	}
}

type Function struct {
	id              int64
	key             NodeKey
	statementsBlock *StatementsBlock
}

func (f *Function) ID() int64        { return f.id }
func (f *Function) NodeKey() NodeKey { return f.key }

func (f *Function) Exec(env *environment.Environment, _ *object.Result, execMngr execManager) error {
	return f.statementsBlock.Exec(env, execMngr)
}
