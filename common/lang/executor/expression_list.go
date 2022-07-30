package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

func NewExpressionList(exprs []Expr) *ExpressionList {
	return &ExpressionList{
		key:   KeyExpressionList,
		exprs: exprs,
	}
}

type ExpressionList struct {
	id    int64
	key   NodeKey
	exprs []Expr
}

func (el *ExpressionList) NodeKey() NodeKey { return el.key }
func (el *ExpressionList) ID() int64        { return el.id }

func (el *ExpressionList) Exec(env *fdalang.Environment, result *Result, execQueue *ExecFnList) error {
	fn := execQueue.Current()
	results := make([]*Result, len(el.exprs))
	for i := range el.exprs {
		ii := i
		results[ii] = NewResult()
		fn = execQueue.AddAfter(fn, el.exprs[ii], func() error {
			return el.exprs[ii].Exec(env, results[ii], execQueue)
		})
	}
	for i := range el.exprs {
		ii := i
		fn = execQueue.AddAfter(fn, el.exprs[ii], func() error {
			result.Add(results[ii].objectList[0])
			return nil
		})
	}
	return nil
}
