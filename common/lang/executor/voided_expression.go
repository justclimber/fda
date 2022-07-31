package executor

func NewVoidedExpression(expr Expr) *VoidedExpression {
	return &VoidedExpression{
		key:  KeyVoidedExpression,
		expr: expr,
	}
}

type VoidedExpression struct {
	id   int64
	key  NodeKey
	expr Expr
}

func (v *VoidedExpression) ID() int64        { return v.id }
func (v *VoidedExpression) NodeKey() NodeKey { return v.key }

func (v *VoidedExpression) Exec(env *Environment, executor execManager) error {
	executor.AddNextExec(v.expr, func() error {
		return v.expr.Exec(env, NewResult(), executor)
	})
	return nil
}
