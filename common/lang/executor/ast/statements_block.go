package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func NewStatementsBlock(id int64, stmts []Stmt) *StatementsBlock {
	return &StatementsBlock{
		id:         id,
		statements: stmts,
	}
}

type StatementsBlock struct {
	id         int64
	statements []Stmt
}

func (sb *StatementsBlock) ID() int64        { return sb.id }
func (sb *StatementsBlock) NodeKey() NodeKey { return KeyStatementsBlock }

func (sb *StatementsBlock) Exec(env *environment.Environment, execMngr execManager) error {
	for _, statement := range sb.statements {
		err := statement.Exec(env, execMngr)
		if err != nil {
			return err
		}
	}
	return nil
}
