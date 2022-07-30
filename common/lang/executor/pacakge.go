package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

func NewPackage(mainFunction *Function) *Package {
	return &Package{
		key:          KeyPackage,
		mainFunction: mainFunction,
	}
}

type Package struct {
	id           int64
	key          NodeKey
	mainFunction *Function
}

func (p *Package) ID() int64        { return p.id }
func (p *Package) NodeKey() NodeKey { return p.key }

func (p *Package) Exec(env *fdalang.Environment, execQueue *ExecFnList) error {
	return p.mainFunction.Exec(env, NewResult(), execQueue)
}
