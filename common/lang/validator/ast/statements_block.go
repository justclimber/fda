package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
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

func (sb *StatementsBlock) ID() int64            { return sb.id }
func (sb *StatementsBlock) NodeKey() ast.NodeKey { return ast.KeyStatementsBlock }

func (sb *StatementsBlock) Check(env ValidatorEnv, validMngr validationManager) (*execAst.StatementsBlock, error) {
	statementsAst := make([]execAst.Stmt, 0, len(sb.statements))
	for _, statement := range sb.statements {
		stmtAst, err := statement.Check(env, validMngr)
		if err != nil {
			return nil, err
		}
		statementsAst = append(statementsAst, stmtAst)
	}
	return execAst.NewStatementsBlock(sb.id, statementsAst), nil
}
