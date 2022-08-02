package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewFunctionCall(name string) *FunctionCall {
	return &FunctionCall{
		key:  KeyFunctionCall,
		name: name,
	}
}

type FunctionCall struct {
	id   int64
	key  NodeKey
	name string
}

func (fc *FunctionCall) ID() int64        { return fc.id }
func (fc *FunctionCall) NodeKey() NodeKey { return fc.key }

func (fc *FunctionCall) Exec(env *environment.Environment, _ *object.Result, execMngr execManager) error {
	// todo compile time check?
	definition, _ := execMngr.MainPackage().FunctionDefinition(fc.name)
	functionEnv := environment.NewEnclosedEnvironment(env)
	if definition.args != nil {
		for _, arg := range definition.args {
			// todo compile time check?
			inputArg, _ := env.Get(arg.varName)
			functionEnv.Set(arg.varName, inputArg)
		}
	}
	execMngr.AddNextExec(definition.statementsBlock, func() error {
		return definition.statementsBlock.Exec(functionEnv, execMngr)
	})
	return nil
}
