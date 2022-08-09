package computer

import (
	"errors"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
	"github.com/justclimber/fda/server/command"
)

type Computer struct {
	executor *executor.Executor
	env      *environment.Environment
	code     *ast.FunctionCall
}

func NewComputer(exec *executor.Executor, env *environment.Environment) *Computer {
	return &Computer{executor: exec, env: env}
}

func (c *Computer) execAll() (*object.Result, error) {
	if c.code == nil {
		return nil, errors.New("no code to exec")
	}
	return c.executor.ExecAll(c.env, c.code)
}

func (c *Computer) Run() (command.Command, error) {
	if c.code == nil {
		return command.Command{}, nil
	}
	res, err := c.execAll()
	if err != nil {
		return command.Command{}, err
	}
	objFloat, _ := res.ObjectList[0].(*object.ObjFloat)
	return command.Command{
		Move: objFloat.Value,
	}, nil
}

func (c *Computer) SetCode(f *ast.FunctionCall) {
	c.code = f
}
