package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewComparisonOperation(left, right Expr, operator object.ComparisonOperator) *ComparisonOperation {
	return &ComparisonOperation{
		key:      KeyComparisonOperation,
		left:     left,
		right:    right,
		operator: operator,
	}
}

type ComparisonOperation struct {
	id       int64
	key      NodeKey
	left     Expr
	right    Expr
	operator object.ComparisonOperator
}

func (ao *ComparisonOperation) ID() int64        { return ao.id }
func (ao *ComparisonOperation) NodeKey() NodeKey { return ao.key }

func (ao *ComparisonOperation) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	res := object.NewResult()
	execMngr.AddNextExec(ao.left, func() error {
		return ao.left.Exec(env, res, execMngr)
	})
	execMngr.AddNextExec(ao.right, func() error {
		return ao.right.Exec(env, res, execMngr)
	})
	execMngr.AddNextExec(ao, func() error {
		switch ao.operator {
		case object.OperatorEqual:
			result.Add(res.DoEqual())
		case object.OperatorNotEqual:
			result.Add(res.DoNotEqual())
		case object.OperatorGraterThan:
			result.Add(res.DoGraterThan())
		case object.OperatorLessThan:
			result.Add(res.DoLessThan())
		case object.OperatorGraterOrEqualThan:
			result.Add(res.DoGraterThanOrEqual())
		case object.OperatorLessOrEqualThan:
			result.Add(res.DoLessThanOrEqual())
		}
		return nil
	})
	return nil
}
