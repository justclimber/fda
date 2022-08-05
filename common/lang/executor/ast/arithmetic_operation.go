package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewArithmeticOperation(left, right Expr, operator object.ArithmeticOperator) *ArithmeticOperation {
	return &ArithmeticOperation{
		left:     left,
		right:    right,
		operator: operator,
	}
}

type ArithmeticOperation struct {
	id       int64
	left     Expr
	right    Expr
	operator object.ArithmeticOperator
}

func (ao *ArithmeticOperation) ID() int64        { return ao.id }
func (ao *ArithmeticOperation) NodeKey() NodeKey { return KeyArithmeticOperation }

func (ao *ArithmeticOperation) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	res := object.NewResult()
	execMngr.AddNextExec(ao.left, func() error {
		return ao.left.Exec(env, res, execMngr)
	})
	execMngr.AddNextExec(ao.right, func() error {
		return ao.right.Exec(env, res, execMngr)
	})
	execMngr.AddNextExec(ao, func() error {
		switch ao.operator {
		case object.OperatorAddition:
			result.Add(res.DoAddition())
		case object.OperatorSubtraction:
			result.Add(res.DoSubtraction())
		case object.OperatorMultiplication:
			result.Add(res.DoMultiplication())
		case object.OperatorDivision:
			result.Add(res.DoDivision())
		}
		return nil
	})
	return nil
}
