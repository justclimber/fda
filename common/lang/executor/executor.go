package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

type Executor struct {
	mainPackage *Package
	env         *fdalang.Environment
	execQueue   *ExecFnList
}

func NewExecutor(env *fdalang.Environment, mainPackage *Package) *Executor {
	return &Executor{
		mainPackage: mainPackage,
		env:         env,
		execQueue:   NewExecFnList(),
	}
}

func (e *Executor) Exec() error {
	err := e.mainPackage.Exec(e.env, e.execQueue)
	if err != nil {
		return err
	}

	hasNext := false
	for {
		hasNext, err = e.Next()
		if err != nil {
			return err
		}
		if !hasNext {
			break
		}
	}
	return nil
}

func (e *Executor) Next() (bool, error) {
	hasNext, err := e.execQueue.Next()
	if err != nil {
		return false, err
	}
	return hasNext, nil
}

func (e *Executor) DebugPrint() {
	e.env.Print()
}

func NewExecFnList() *ExecFnList {
	return &ExecFnList{}
}

type ExecFnList struct {
	head *ExecFn
	curr *ExecFn
}

type ExecFn struct {
	fn   func() error
	next *ExecFn
	node Node
}

func (el *ExecFnList) Next() (bool, error) {
	if el.curr == nil {
		return false, nil
	}
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

func (el *ExecFnList) AddAfterCurrent(node Node, fn func() error) *ExecFn {
	execFn := &ExecFn{fn: fn, node: node}
	if el.head == nil {
		el.head = execFn
		el.curr = execFn
		return execFn
	}
	afterCurr := el.curr.next
	execFn.next = afterCurr
	el.curr.next = execFn
	return execFn
}

func (el *ExecFnList) Current() *ExecFn {
	return el.curr
}

func (el *ExecFnList) AddAfter(execFn *ExecFn, node Node, fn func() error) *ExecFn {
	newExecFn := &ExecFn{fn: fn, node: node}

	afterCurr := execFn.next
	newExecFn.next = afterCurr
	execFn.next = newExecFn
	return newExecFn
}
