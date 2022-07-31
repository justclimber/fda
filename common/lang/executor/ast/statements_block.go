package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func NewStatementsBlock(stmts []Stmt) *StatementsBlock {
	return &StatementsBlock{
		key:        KeyStatementsBlock,
		statements: stmts,
	}
}

type StatementsBlock struct {
	id         int64
	key        NodeKey
	statements []Stmt
}

func (sb *StatementsBlock) ID() int64        { return sb.id }
func (sb *StatementsBlock) NodeKey() NodeKey { return sb.key }

func (sb *StatementsBlock) Exec(env *environment.Environment, execMngr execManager) error {
	for _, statement := range sb.statements {
		err := statement.Exec(env, execMngr)
		if err != nil {
			return err
		}
	}
	return nil
}
