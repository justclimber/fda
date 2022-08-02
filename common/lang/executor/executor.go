package executor

import (
	"github.com/justclimber/fda/common/lang/executor/ast"
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

func (e *Executor) Exec(env *environment.Environment, function *ast.FunctionCall) error {
	err := function.Exec(env, object.NewResult(), e)
	if err != nil {
		return err
	}

	hasNext := false
	for {
		hasNext, err = e.ExecNext()
		if err != nil {
			return err
		}
		if !hasNext {
			break
		}
	}
	return nil
}

func (e *Executor) ExecNext() (bool, error) {
	hasNext, err := e.execQueue.ExecNext()
	if err != nil {
		return false, err
	}
	return hasNext, nil
}

func (e *Executor) AddNextExec(node ast.Node, fn func() error) {
	e.execQueue.AddNext(node, fn)
}

func (e *Executor) MainPackage() *ast.Package {
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
