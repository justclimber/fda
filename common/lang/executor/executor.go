package executor

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type Executor struct {
	packagist *Packagist
	execQueue *ExecFnList
}

func NewExecutor(packagist *Packagist, execQueue *ExecFnList) *Executor {
	return &Executor{
		packagist: packagist,
		execQueue: execQueue,
	}
}

func (e *Executor) ExecAll(env *environment.Environment, function *execAst.FunctionCall) (*object.Result, error) {
	result := object.NewResult()
	err := function.Exec(env, result, e)
	if err != nil {
		return nil, err
	}

	hasNext := false
	for {
		hasNext, err = e.ExecNext()
		if err != nil {
			return nil, err
		}
		if !hasNext {
			break
		}
	}
	return result, nil
}

func (e *Executor) ExecNext() (bool, error) {
	return e.execQueue.ExecNext()
}

func (e *Executor) AddNextExec(node ast.Node, fn func() error) {
	e.execQueue.AddNext(node, fn)
}

func (e *Executor) MainPackage() *execAst.Package {
	return e.packagist.Main()
}

func NewExecFnList() *ExecFnList {
	return &ExecFnList{}
}

type ExecFnList struct {
	head *ExecFn
	curr *ExecFn
	next *ExecFn
}

type ExecFn struct {
	fn   func() error
	next *ExecFn
	node ast.Node
}

func (el *ExecFnList) ExecNext() (bool, error) {
	if el.curr == nil {
		return false, nil
	}
	el.next = el.curr
	err := el.curr.fn()
	if err != nil {
		return false, err
	}
	if el.curr.next == nil {
		return false, nil
	}
	el.curr = el.curr.next
	return true, nil
}

func (el *ExecFnList) AddNext(node ast.Node, fn func() error) {
	newExecFn := &ExecFn{fn: fn, node: node}

	if el.head == nil {
		el.head = newExecFn
		el.curr = newExecFn
		el.next = newExecFn
		return
	}

	afterNext := el.next.next
	newExecFn.next = afterNext
	el.next.next = newExecFn
	el.next = newExecFn
}
